
import '../css/index.scss'
import 'react-select/dist/react-select.css';

import 'babel-polyfill'
import React from 'react'
import ReactDOM from 'react-dom'
import { createStore, applyMiddleware, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import thunkMiddleware from 'redux-thunk'
import multiMiddleware from 'redux-multi'
import { batchedSubscribe } from 'redux-batched-subscribe'

import * as actions from './actions'
import expeditionReducer from './reducers/expeditions'
import authReducer from './reducers/auth'
import { Router, Route, IndexRoute, Redirect, browserHistory } from 'react-router'

import Root from './components/Root'
import ProfileSection from './components/AdminPage/ProfileSection'
import LandingPageContainer from './containers/LandingPage/LandingPage'
import AdminPageContainer from './containers/AdminPage/AdminPage'
import NewGeneralSettingsContainer from './containers/AdminPage/NewExpeditionPage/NewExpeditionPage'
import NewInputsContainer from './containers/AdminPage/NewExpeditionPage/InputsSection'
import NewConfirmationContainer from './containers/AdminPage/NewExpeditionPage/ConfirmationSection'
import ExpeditionPageContainer from './containers/AdminPage/ExpeditionPage/ExpeditionPage'
import GeneralSettingsContainer from './containers/AdminPage/ExpeditionPage/GeneralSettingsSection'
import InputsContainer from './containers/AdminPage/ExpeditionPage/InputsSection'
import TeamsContainer from './containers/AdminPage/ExpeditionPage/TeamsSection'
import UploaderContainer from './containers/AdminPage/ExpeditionPage/UploaderSection'
import ThemeContainer from './containers/AdminPage/ExpeditionPage/ThemeSection'

import ProjectGridContainer from './containers/AdminPage/ProjectGrid'
import ProjectSettingsContainer from './containers/AdminPage/ProjectSettings'


import FKApiClient from './api/api.js'

document.getElementById('root').remove()

const createStoreWithMiddleware = applyMiddleware(
  thunkMiddleware,
  multiMiddleware,
)(createStore)
const reducer = combineReducers({
  auth: authReducer,
  expeditions: expeditionReducer,
})
const store = createStoreWithMiddleware(reducer)


function checkAuthentication(state, replace) {  
  if (!FKApiClient.signedIn()) {
    replace({
      pathname: '/signin',
      state: { nextPathname: state.location.pathname }
    })
  } 
}

const routes = (
  <Route path="/" component={Root}>
    <IndexRoute component={LandingPageContainer}/>
    <Route component={LandingPageContainer}>
      <Route path="signup"/>
      <Route path="signin"/>
    </Route>

    <Route 
      path="admin"
      component={ AdminPageContainer }
      onEnter={(state, replace, callback) => {
        checkAuthentication(state, replace)
        store.dispatch(actions.requestUser((user) => {
          store.dispatch(actions.requestProjects((projects) => {
            callback()
          }))
        }))
      }}
    >
      <IndexRoute
        component={ ProjectGridContainer }
        onEnter={(state) => {
          store.dispatch(actions.setBreadcrumbs(0, 'Projects', '/admin'))
        }}
      />
      <Route 
        path="profile"
        onEnter={(state) => {
          store.dispatch(actions.setBreadcrumbs(0, 'Profile settings', '/profile'))
        }}
      />
      <Route
        path=":projectID"
        onEnter={(state, replace, callback) => {
          const projectID = state.params.projectID
          const projectName = store.getState().expeditions.getIn(['projects', projectID, 'name'])
          let expeditionID = state.params.expeditionID
          store.dispatch(actions.setCurrentProject(projectID))
          store.dispatch(actions.setBreadcrumbs(0, 'Project: ' + projectName, '/admin/' + projectID))
          store.dispatch(actions.requestExpeditions(projectID, (expeditions) => {
            callback()
          }))
        }}
        onChange={(state) => {
          const projectID = state.params.projectID
          const projectName = store.getState().expeditions.getIn(['projects', projectID, 'name'])
          store.dispatch(actions.setBreadcrumbs(0, 'Project: ' + projectName, '/admin/' + projectID))
        }}
      >
        <IndexRoute
          component={ ProjectSettingsContainer }
        />
        <Route
          path=":expeditionID"
          component={ ExpeditionPageContainer }
          onLeave={() => {
            store.dispatch(actions.setCurrentExpedition(null))
          }}
          onEnter={(state) => {
            const expeditionID = state.params.expeditionID
            const expeditionName = store.getState().expeditions.getIn(['expeditions', expeditionID, 'name'])
            store.dispatch(actions.setCurrentExpedition(expeditionID))
            store.dispatch(actions.setBreadcrumbs(1, 'Expedition: ' + expeditionName, '/admin/' + state.params.projectID + '/' + expeditionID))
          }}
          onChange={(state) => {
            const expeditionID = state.params.expeditionID
            const expeditionName = store.getState().expeditions.getIn(['expeditions', expeditionID, 'name'])
            store.dispatch(actions.setBreadcrumbs(1, 'Expedition: ' + expeditionName, '/admin/' + state.params.projectID + '/' + expeditionID))
          }}
        >
          <IndexRoute component={ GeneralSettingsContainer }
            onEnter={(state) => {
              store.dispatch(actions.setBreadcrumbs(2, 'General Settings', '/admin/' + state.params.projectID + '/' + state.params.expeditionID + '/general-settings'))
            }}
          />
          <Route
            path="inputs"
            component={InputsContainer}
            onEnter={(state, replace, callback) => {
              store.dispatch(actions.initInputsPage(() => {
                store.dispatch(actions.setBreadcrumbs(2, 'Inputs Settings', '/admin/' + state.params.projectID + '/' + state.params.expeditionID + '/inputs'))
                callback()
              }))
            }}
          />
          <Route
            path="teams"
            component={TeamsContainer}
            onEnter={(state, replace, callback) => {
              store.dispatch(actions.initTeamsPage(() => {
                store.dispatch(actions.setBreadcrumbs(2, 'Teams Settings', '/admin/' + state.params.projectID + '/' + state.params.expeditionID + '/teams'))
                callback()
              }))
            }}
          />
          <Route
            path="uploader"
            component={UploaderContainer}
            onEnter={(state, replace, callback) => {
              store.dispatch(actions.initUploaderPage(() => {
                store.dispatch(actions.setBreadcrumbs(2, 'Data Uploader', '/admin/' + state.params.projectID + '/' + state.params.expeditionID + '/uploader'))
                callback()
              }))
            }}
          />
          <Route
            path="theme"
            component={ThemeContainer}
            onEnter={(state, replace, callback) => {
              store.dispatch(actions.initThemePage(() => {
                store.dispatch(actions.setBreadcrumbs(2, 'Theme Settings', '/admin/' + state.params.projectID + '/' + state.params.expeditionID + '/theme'))
                callback()
              }))
            }}
          />
        </Route>
      </Route>
    </Route>
  </Route>
)

var render = function () {
  ReactDOM.render(
    (
      <Provider store={store}>
        <Router history={browserHistory} routes={routes}/>
      </Provider>
    ),
    document.getElementById('fieldkit')
  )
}

render()