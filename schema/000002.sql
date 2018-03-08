DROP FUNCTION IF EXISTS fieldkit.fk_clustered_identical(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_tracks(desired_source_id BIGINT);

DROP FUNCTION IF EXISTS fieldkit.fk_spatial_clusters(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_temporal_clusters(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_temporal_geometries(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_source_summary(desired_source_id BIGINT);

DROP FUNCTION IF EXISTS fieldkit.fk_update_spatial_clusters(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_update_temporal_clusters(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_update_temporal_geometries(desired_source_id BIGINT);
DROP FUNCTION IF EXISTS fieldkit.fk_update_source_summary(desired_source_id BIGINT);

CREATE TABLE fieldkit.sources_summaries (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    updated_at timestamp NOT NULL,
    number_of_features integer NOT NULL,
    last_feature_id integer NOT NULL,
    start_time timestamp NOT NULL,
    end_time timestamp NOT NULL,
    centroid geometry(POINT, 4326) NOT NULL,
    radius decimal NOT NULL,
    PRIMARY KEY (source_id)
);

CREATE TABLE fieldkit.sources_temporal_clusters (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    number_of_features integer NOT NULL,
    start_time timestamp NOT NULL,
    end_time timestamp NOT NULL,
    centroid geometry(POINT, 4326) NOT NULL,
    radius decimal NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);

CREATE TABLE fieldkit.sources_temporal_geometries (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    geometry geometry(LINESTRING, 4326) NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);

CREATE TABLE fieldkit.sources_spatial_clusters (
    source_id integer REFERENCES fieldkit.input (id) ON DELETE CASCADE NOT NULL,
    cluster_id integer NOT NULL,
    updated_at timestamp NOT NULL,
    number_of_features integer NOT NULL,
    start_time timestamp NOT NULL,
    end_time timestamp NOT NULL,
    centroid geometry(POINT, 4326) NOT NULL,
    bounding_circle geometry NOT NULL,
    bounding_envelope geometry NOT NULL,
    radius decimal NOT NULL,
    PRIMARY KEY (source_id, cluster_id)
);

CREATE OR REPLACE FUNCTION fieldkit.fk_clustered_identical(desired_source_id BIGINT)
RETURNS TABLE (
	"source_id" INTEGER,
	"location" geometry,
	"size" BIGINT,
  "min_timestamp" timestamp,
  "max_timestamp" timestamp,
	"copy" BIGINT,
	"spatial_cluster_id" integer
) AS
'
BEGIN
RETURN QUERY
WITH
  with_identical_clustering AS (
    SELECT
      d.input_id AS source_id,
      d.location,
      COUNT(*) AS actual_size,
      MIN(d.timestamp) AS min_timestamp,
      MAX(d.timestamp) AS max_timestamp,
      LEAST(CAST(11 AS BIGINT), COUNT(*)) AS capped_size
    FROM fieldkit.document d
    WHERE d.input_id IN (desired_source_id) AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0
    GROUP BY d.input_id, d.location
  ),
  with_cluster_ids AS (
    SELECT
      s.source_id,
      s.location,
      s.actual_size,
      s.min_timestamp,
      s.max_timestamp,
      n AS n,
      ST_ClusterDBSCAN(ST_Transform(s.location, 2950), eps := 50, minPoints := 10) OVER () AS spatial_cluster_id
    FROM with_identical_clustering s, generate_series(0, s.capped_size - 1) AS x(n)
  )
  SELECT *
  FROM with_cluster_ids s
  WHERE s.spatial_cluster_id IS NOT NULL;
END
' LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION fieldkit.fk_tracks(desired_source_id BIGINT)
RETURNS TABLE (
	"source_id" BIGINT,
	"timestamp" timestamp,
	"min_timestamp" timestamp,
	"max_timestamp" timestamp,
	"time_difference" float,
  "actual_size" bigint,
	"temporal_cluster_id" BIGINT,
	"location" geometry
) AS
'
BEGIN
RETURN QUERY
WITH
source AS (
	SELECT
    to_timestamp(AVG(extract(epoch from d.timestamp)))::timestamp AS timestamp,
		MIN(d.timestamp) AS min_timestamp,
		MAX(d.timestamp) AS max_timestamp,
		MAX(d.timestamp) - MIN(d.timestamp) AS min_max_diff,
		d.location,
    (SELECT bool_or(ST_DWithin(d.location::geography, spatial.bounding_envelope::geography, 50)) FROM fieldkit.fk_spatial_clusters(desired_source_id) spatial WHERE spatial.source_id = 125) AS in_spatial,
    COUNT(d.location) AS actual_size
	FROM fieldkit.document d
  WHERE d.input_id IN (125) AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0
  GROUP BY d.location
),
with_timestamp_differences AS (
	SELECT
		*,
  			                              LAG(s.timestamp) OVER (ORDER BY s.timestamp) AS previous_timestamp,
		EXTRACT(epoch FROM (s.timestamp - LAG(s.timestamp) OVER (ORDER BY s.timestamp))) AS time_difference
	FROM source s
  WHERE NOT s.in_spatial
	ORDER BY s.timestamp
),
with_temporal_clustering AS (
	SELECT
		*,
		CASE WHEN s.time_difference > 600
			OR s.time_difference IS NULL THEN true
			ELSE NULL
		END AS new_temporal_cluster
	FROM with_timestamp_differences s
),
with_assigned_temporal_clustering AS (
	SELECT
		*,
		COUNT(new_temporal_cluster/* OR spatial_cluster_change*/) OVER (
			ORDER BY s.timestamp
			ROWS UNBOUNDED PRECEDING
		) AS temporal_cluster_id
	FROM with_temporal_clustering s
)
SELECT
	desired_source_id AS source_id,
  s.timestamp,
  s.min_timestamp,
  s.max_timestamp,
  s.time_difference,
  s.actual_size,
  s.temporal_cluster_id,
  s.location
FROM with_assigned_temporal_clustering s;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_spatial_clusters(desired_source_id BIGINT)
RETURNS SETOF fieldkit.sources_spatial_clusters
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_source_id::integer AS source_id,
  f.spatial_cluster_id,
  NOW()::timestamp,
  SUM(CASE copy WHEN 0 THEN f.size ELSE 0 END)::integer,
  MIN(min_timestamp),
  MAX(max_timestamp),
  ST_Centroid(ST_Collect(f.location))::geometry(Point, 4326) AS centroid,
  ST_MinimumBoundingCircle(ST_Collect(f.location)) AS bounding_circle,
  ST_Envelope(ST_Collect(f.location)) AS bounding_envelope,
  SQRT(ST_Area(ST_MinimumBoundingCircle(ST_Collect(f.location))::geography) / pi())::numeric AS radius
FROM fieldkit.fk_clustered_identical(desired_source_id) AS f
GROUP BY spatial_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_temporal_clusters(desired_source_id BIGINT)
RETURNS SETOF fieldkit.sources_temporal_clusters
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_source_id::integer AS source_id,
  f.temporal_cluster_id::integer,
  NOW()::timestamp,
  SUM(actual_size)::integer,
  MIN(min_timestamp),
  MAX(max_timestamp),
  ST_Centroid(ST_Collect(f.location))::geometry(Point, 4326) AS centroid,
  SQRT(ST_Area(ST_MinimumBoundingCircle(ST_Collect(f.location))::geography) / pi())::numeric AS radius
FROM fieldkit.fk_tracks(desired_source_id) AS f
GROUP BY temporal_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_temporal_geometries(desired_source_id BIGINT)
RETURNS SETOF fieldkit.sources_temporal_geometries
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_source_id::integer AS source_id,
  f.temporal_cluster_id::integer,
  NOW()::timestamp,
  ST_LineFromMultiPoint(ST_Collect(f.location))::geometry(LineString, 4326)
FROM fieldkit.fk_tracks(desired_source_id) AS f
GROUP BY temporal_cluster_id;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_source_summary(desired_source_id BIGINT)
RETURNS SETOF fieldkit.sources_summaries
AS
'
BEGIN
RETURN QUERY
SELECT
  desired_source_id::integer AS source_id,
  NOW()::timestamp AS updated_at,
	(SELECT COUNT(d.id)::integer AS number_of_features FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT d.Id::integer AS last_feature_id FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0 ORDER BY d.timestamp DESC LIMIT 1),
	(SELECT MIN(d.timestamp) AS start_time FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT MAX(d.timestamp) AS end_time FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT ST_Centroid(ST_Collect(d.location))::geometry(Point, 4326) AS centroid FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0),
	(SELECT Sqrt(ST_Area(ST_MinimumBoundingCircle(ST_Collect(d.location))::geography))::numeric AS radius FROM fieldkit.document AS d WHERE d.visible AND d.input_id = desired_source_id AND ST_X(d.location) != 0 AND ST_Y(d.location) != 0);
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_source_summary(desired_source_id BIGINT)
RETURNS SETOF integer
AS
'
BEGIN
  INSERT INTO fieldkit.sources_summaries
  SELECT * FROM fieldkit.fk_source_summary(desired_source_id)
	ON CONFLICT (source_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    last_feature_id = excluded.last_feature_id,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    centroid = excluded.centroid,
    radius = excluded.radius;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_spatial_clusters(desired_source_id BIGINT)
RETURNS SETOF integer
AS
'
BEGIN
  DELETE FROM fieldkit.sources_spatial_clusters WHERE (source_id = desired_source_id);
  INSERT INTO fieldkit.sources_spatial_clusters
  SELECT * FROM fieldkit.fk_spatial_clusters(desired_source_id)
  ON CONFLICT (source_id, cluster_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    centroid = excluded.centroid,
    radius = excluded.radius;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_temporal_clusters(desired_source_id BIGINT)
RETURNS SETOF integer
AS
'
BEGIN
  DELETE FROM fieldkit.sources_temporal_clusters WHERE (source_id = desired_source_id);
  INSERT INTO fieldkit.sources_temporal_clusters
  SELECT * FROM fieldkit.fk_temporal_clusters(desired_source_id)
  ON CONFLICT (source_id, cluster_id) DO UPDATE SET
    updated_at = excluded.updated_at,
    number_of_features = excluded.number_of_features,
    start_time = excluded.start_time,
    end_time = excluded.end_time,
    centroid = excluded.centroid,
    radius = excluded.radius;
END
' LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fieldkit.fk_update_temporal_geometries(desired_source_id BIGINT)
RETURNS SETOF integer
AS
'
BEGIN
  DELETE FROM fieldkit.sources_temporal_geometries WHERE (source_id = desired_source_id);
  INSERT INTO fieldkit.sources_temporal_geometries
  SELECT * FROM fieldkit.fk_temporal_geometries(desired_source_id)
  ON CONFLICT (source_id, cluster_id) DO UPDATE SET
     updated_at = excluded.updated_at,
     geometry = excluded.geometry;
END
' LANGUAGE plpgsql;
