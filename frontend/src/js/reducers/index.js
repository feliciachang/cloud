import _ from 'lodash';

import { combineReducers } from 'redux';

import * as ActionTypes from '../actions/types';

import { FkGeoJSON } from '../common/geojson';

function activeExpedition(state = { project: null, expedition: null }, action) {
    switch (action.type) {
    case ActionTypes.API_PROJECT_GET.SUCCESS:
        return Object.assign({ }, state, { project: action.response });
    case ActionTypes.API_EXPEDITION_GET.SUCCESS:
        return Object.assign({ }, state, { expedition: action.response });
    default:
        return state;
    }
}

const visibleFeaturesInitialState = {
    focus: { features: [] },
    sources: { }
};

function visibleFeatures(state = visibleFeaturesInitialState, action) {
    switch (action.type) {
    case ActionTypes.API_EXPEDITION_GEOJSON_GET.SUCCESS: {
        let newGeojson = Object.assign({}, action.response.geo);
        if (state.geojson) {
            newGeojson.features = [ ...state.geojson.features, ...newGeojson.features ]
        }
        return Object.assign({ }, state, {
            geojson: newGeojson
        });
    }
    case ActionTypes.FOCUS_FEATURE:
        return Object.assign({ }, state, {
            focus: {
                feature: action.feature,
                center: action.feature.geometry.coordinates,
                features: []
            }
        });
    case ActionTypes.API_FEATURE_GEOJSON_GET.SUCCESS: {
        const feature = action.response.geo.features[0];
        const nextState = Object.assign({}, state);
        _.each(nextState.sources, (source, id) => {
            if (source.lastFeatureId === feature.properties.id) {
                source.lastFeature = feature;
            }
        });
        return nextState;
    }
    case ActionTypes.API_SOURCE_GET.SUCCESS: {
        const source = action.response;
        const nextState = Object.assign({}, state);
        nextState.sources[source.id] = {...source, ...{ } }
        return nextState;
    }
    case ActionTypes.FOCUS_TIME:
        const expedition = new FkGeoJSON(state.geojson);
        const features = expedition.getFeaturesWithinTime(action.time, 500000 * 10);

        return Object.assign({ }, state, {
            focus: {
                time: action.time,
                center: action.center,
                features: features
            }
        });
    default:
        return state;
    }
}

function playbackMode(state = { }, action) {
    switch (action.type) {
    case ActionTypes.CHANGE_PLAYBACK_MODE:
        return action.mode; // TODO: Treat this object as immutable. Needs better protection.
    default:
        return state;
    }
}

export default combineReducers({
    activeExpedition,
    visibleFeatures,
    playbackMode,
});
