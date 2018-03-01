// @flow weak

import React, { Component } from 'react';
import ReactMapboxGl, { ScaleControl, ZoomControl/*, Popup*/ } from 'react-mapbox-gl';
import _ from 'lodash';

import { MAPBOX_ACCESS_TOKEN, MAPBOX_STYLE } from '../../secrets';

import BubbleMap from '../visualizations/BubbleMap';
import ClusterMap  from '../visualizations/ClusterMap';
import PlaybackControl from '../PlaybackControl';

import type { Coordinates, Bounds, GeoJSONFeature, GeoJSON } from '../../types/MapTypes';
import type { Focus } from '../../types';

import { FieldKitLogo } from '../icons/Icons';

import FeaturePanel from './FeaturePanel';
import NotificationsPanel from './NotificationsPanel';
import ChartComponent from '../ChartComponent';
import FiltersPanel from './FiltersPanel';

function getFitBounds(geojson: GeoJSON) {
    const lon = geojson.features.map(f => f.geometry.coordinates[0]);
    const lat = geojson.features.map(f => f.geometry.coordinates[1]);
    return [[Math.min(...lon), Math.min(...lat)], [Math.max(...lon), Math.max(...lat)]];
}

function getCenter(fitBounds: Bounds) {
    return [(fitBounds[0][0] + fitBounds[1][0]) / 2, (fitBounds[0][1] + fitBounds[1][1]) / 2];
}

type Props = {
    playbackMode: mixed,
    focusSource: () => mixed,
    onChangePlaybackMode: () => mixed,
    notifyOfUserMapActivity: () => mixed,
    pointDecorator: PointDecorator,
    visibleFeatures: {
        geojson: GeoJSON,
        focus: Focus,
    }
};

const panelContainerStyle = {
    backgroundColor: '#f9f9f9',
    color: "#000",
    position: 'absolute',
    zIndex: 10,
    display: 'flex',
    flexDirection: 'column',
    boxShadow: '0px 1px 4px rgba(0, 0, 0, .3)',
};

const panelHeaderStyle = {
    padding: '5px',
    backgroundColor: 'rgb(196, 196, 196)',
    borderBottom: '1px solid rgb(128, 128, 128)',
    fontWeight: 'bold',
};

const panelBodyStyle = {
};

export class MapBottom extends Component {
    positionStyle() {
        const { sidePanelVisible } = this.props;
        if (!sidePanelVisible) {
            return { top: 'auto', right: 0, bottom: 0, left: 0, height: '250px', width: '100%' };
        }
        return { top: 'auto', right: 300, bottom: 0, left: 0, height: '250px' };
    }

    render() {
        const { onHide, children } = this.props;
        const position = this.positionStyle();

        return (
            <div style={{...panelContainerStyle, ...position}}>
                <div style={{ ...panelHeaderStyle }}><span style={{ cursor: 'pointer' }} onClick={() => onHide()}>Close</span></div>
                <div style={{ ...panelBodyStyle }}>
                    {children}
                </div>
            </div>
        );
    }
}

export class MapRight extends Component {
    render() {
        const { onHide, children } = this.props;
        const position = { top: 0, right: 0, bottom: 'auto', left: 'auto', width: '300px', height: '100%' };

        return (
            <div style={{...panelContainerStyle, ...position}}>
                <div style={{ ...panelHeaderStyle }}><span style={{ cursor: 'pointer' }} onClick={() => onHide()}>Close</span></div>
                <div style={{ ...panelBodyStyle }}>
                    {children}
                </div>
            </div>
        );
    }
}

const Map = ReactMapboxGl({
    accessToken: MAPBOX_ACCESS_TOKEN
});

export default class MapContainer extends Component {
    props: Props
    state: {
        fitBounds: Bounds,
        center: Coordinates,
        zoom: [number],
        feature: ?GeoJSONFeature,
    };

    constructor(props: Props) {
        super(props);
        this.state = {
            fitBounds: null,
            center: null,
            zoom: [14],
            feature: null,
            panels: {
                sidePanelVisible: false,
                bottomPanelVisible: false,
            }
        };
        this.onUserActivityThrottled = _.throttle(this.onUserActivity.bind(this), 500, { leading: true });
    }

    componentWillMount() {
        const { visibleFeatures } = this.props;

        if (false)
        if (visibleFeatures.geojson.features.length > 0) {
            const fitBounds = getFitBounds(visibleFeatures.geojson);
            const center = getCenter(fitBounds);
            this.setState({
                fitBounds,
                center
            });
        }
    }

    componentWillReceiveProps(nextProps) {
        const { visibleFeatures: oldData } = this.props;
        const { visibleFeatures: newData } = nextProps;

        if (oldData !== newData) {
            const { focus: oldFocus } = oldData;
            const { focus: newFocus } = newData;

            if (oldFocus !== newFocus) {
                if (newFocus.feature) {
                    this.setState({
                        center: newFocus.feature.geometry.coordinates
                    });
                }
                else if (newFocus.center) {
                    this.setState({
                        center: newFocus.center
                    });
                }
            }
        }
    }

    onMarkerClick(feature) {
        this.setState({
            feature,
            panels: {
                sidePanelVisible: true,
            },
        });
    }

    onFocusFeature(feature) {
        const { focusFeature } = this.props;
        focusFeature(feature);

        this.setState({
            feature,
            panels: {
                sidePanelVisible: true,
            },
        });
    }

    onShowChart(chart) {
        this.setState({
            chart,
            panels: {
                sidePanelVisible: true,
                bottomPanelVisible: true,
            },
        });
    }

    onFocusSource(source) {
        const { focusSource } = this.props;

        focusSource(source);
    }

    onUserActivity(how) {
        const { notifyOfUserMapActivity } = this.props;

        notifyOfUserMapActivity();
    }

    onPopupChange() {
    }

    renderPanels() {
        const { visibleFeatures, playbackMode } = this.props;
        const { panels, chart, feature } = this.state;

        const { mode } = playbackMode;
        const paused = mode === "Pause";

        return (
            <div>
                { !paused && <NotificationsPanel features={visibleFeatures.focus.features} sidePanelVisible={panels.sidePanelVisible} />}
                { panels.sidePanelVisible && <MapRight onHide={() => this.setState({ panels: { ...panels, sidePanelVisible: false, bottomPanelVisible: false }})}>
                  <FeaturePanel feature={feature} onShowChart={chart => this.onShowChart(chart)}/>
                </MapRight> }
                { panels.bottomPanelVisible && <MapBottom onHide={() => this.setState({ panels: { ...panels, bottomPanelVisible: false }})} sidePanelVisible={panels.sidePanelVisible}>
                  <ChartComponent chart={chart} />
                </MapBottom> }
            </div>
        );
    }

    render() {
        const { pointDecorator, visibleFeatures, onChangePlaybackMode, playbackMode } = this.props;
        const { fitBounds, center, zoom } = this.state;

        if (!center) {
            return <div></div>;
        }

        return (
            <div>
                <Map style={MAPBOX_STYLE} containerStyle={ { height: '100vh', width: '100vw', } }
                    movingMethod="easeTo" center={ center } zoom={ zoom } fitBounds={ fitBounds }
                    onDrag={ this.onUserActivityThrottled }
                    onClick={ this.onPopupChange.bind(this) } containerStyle={ { height: '100vh', width: '100vw', } }>

                    <ClusterMap visibleFeatures={ visibleFeatures } data={ visibleFeatures.sources } click={ this.onMarkerClick.bind(this) } />

                    <BubbleMap pointDecorator={ pointDecorator } data={ visibleFeatures.geojson } click={ this.onMarkerClick.bind(this) } />

                    <ScaleControl style={ { backgroundColor: 'rgba(0, 0, 0, 0)', left: '12px', bottom: '6px', } } />
                    { false && <ZoomControl className="zoom-control" position={ 'topLeft' } />}

                    <PlaybackControl className="playback-control" playback={ playbackMode } onPlaybackChange={ onChangePlaybackMode.bind(this) } />

                    <FiltersPanel features visibleFeatures={visibleFeatures} onShowSource={ this.onFocusSource.bind(this) } onShowFeature={ this.onFocusFeature.bind(this) } />
                </Map>
                {this.renderPanels()}
                <div className="disclaimer-panel">
                    <div className="disclaimer-body">
                        <span className="b">NOTE: </span> Map images have been obtained from a third-party and do not reflect the editorial decisions of National Geographic.
                    </div>
                    <div className="fieldkit-logo">
                        <FieldKitLogo />
                    </div>
                </div>
            </div>
        );
    }
}
