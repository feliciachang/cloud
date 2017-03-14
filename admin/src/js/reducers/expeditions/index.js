
import * as actions from '../../actions'
import I from 'immutable'
import slug from 'slug'

export const initialState = I.fromJS({
  user: null,
  breadcrumbs: [null, null, null],
  errors: null,
  suggestedMembers: null,
  modal: {
    type: null
  },
  currentProject: null,
  currentExpedition: null,
  editedTeam: null,
  projects: null,
  expeditions: null,
  inputs: {
    'conservifymodule': {
      id: 'conservifymodule',
      type: 'sensor',
      name: 'Conservify Module',
      description: 'Lorem ipsum dolor sit amet, consectetur adipiscing.',
      setupType: 'token'
    }
  },
  uploaders: {}
})

const expeditionReducer = (state = initialState, action) => {

  console.log('reducer:', action.type, action)
  // try {
    switch (action.type) {
      case actions.RECEIVE_USER: {
        return state.set('user', action.user)
      }

      case actions.SET_BREADCRUMBS: {
        let newState = state
          .setIn(['breadcrumbs', action.level], action.value)
        for (let i = action.level + 1; i < 3; i ++) {
          newState = newState.setIn(['breadcrumbs', i], null)
        }
        return newState
      }

      case actions.SET_ERROR: {
        return state
          .set('errors', action.errors)
      }

      case actions.PROMPT_MODAL: {      
        return state
          .setIn(['modal', 'type'], action.modalType)
      }

      case actions.CLOSE_MODAL: {
        return state
          .setIn(['modal', 'type'], null)
      }

      case actions.RECEIVE_PROJECTS: {
        return state
          .set('projects', action.projects)
      }

      case actions.RECEIVE_EXPEDITIONS: {
        return state
          .set('expeditions', action.expeditions)
          .setIn(['currentProject', 'expeditions'], action.expeditions.map(e => e.get('id')).toList())
      }

      case actions.RECEIVE_TOKEN: {
        return state
          .setIn(['currentExpedition', 'token'], action.token)
      }

      case actions.NEW_PROJECT: {
        const projectID = 'project-' + Date.now()
        return state
          .set(
            'currentProject', 
            I.fromJS({
              id: projectID,
              name: 'Project Name',
              description: 'Project Description',
              expeditions: []
            })
          )
      }

      case actions.SET_PROJECT_PROPERTY: {
        const newState = state.setIn(
          ['currentProject'].concat(action.keyPath),
          action.value
        )
        if (!!action.keyPath && action.keyPath.length === 1 && action.keyPath[0] === 'name') {
          return newState
            .setIn(['currentProject', 'id'], slug(action.value))
        } else {
          return newState 
        }
      }

      case actions.CANCEL_PROJECT: {
        return state
          .set('currentProject', null)
      }

      case actions.SAVE_PROJECT: {
        return state
          .setIn(['currentProject', 'id'], action.id)
          .setIn(
            ['projects', action.id],
            state.get('currentProject')
              .set('id', action.id)
          )
      }

      case actions.NEW_EXPEDITION: {
        const expeditionID = 'expedition-' + Date.now()
        return state
          .set(
            'currentExpedition', 
            I.fromJS({
              id: expeditionID,
              name: 'Expedition Name',
              description: 'Enter a description',
              startDate: new Date(),
              teams: [],
              inputs: [],
              selectedPreset: null,
              token: ''
            })
          )
      }

      case actions.SET_EXPEDITION_PROPERTY: {
        const newState = state.setIn(
          ['currentExpedition'].concat(action.keyPath),
          action.value
        )
        if (!!action.keyPath && action.keyPath.length === 1 && action.keyPath[0] === 'name') {
          return newState
            .setIn(['currentExpedition', 'id'], slug(action.value))
        } else {
          return newState 
        }
      }

      case actions.SAVE_EXPEDITION: {
        const expedition = state.get('currentExpedition')
        return state.setIn(['expeditions', expedition.get('id')], expedition)
          .set('currentExpedition', null)
      }

      case actions.SET_CURRENT_PROJECT: {
        return state.set('currentProject', state.getIn(['projects', action.projectID]))
      }

      case actions.SET_CURRENT_EXPEDITION: {
        const expedition = state.getIn(['expeditions', action.expeditionID])
        return state
          .set('currentExpedition', expedition)
      }

      case actions.ADD_INPUT: {
        if (!state.getIn(['currentExpedition', 'inputs']).includes(action.id)) {
          return state
            .setIn(
              ['currentExpedition', 'inputs'],
              state.getIn(['currentExpedition', 'inputs']).push(action.id)
            )
            .setIn(
              ['inputs', action.id, 'token'],
              action.token
            )
        } else {
          return state
            .setIn(
              ['inputs', action.id, 'token'],
              action.token
            )
        }
      }

      case actions.REMOVE_INPUT: {
        return state
          .deleteIn(['currentExpedition', 'inputs', action.id])
      }

      case actions.RECEIVE_INPUTS: {
        return state
          .setIn(['currentExpedition', 'inputs'], action.inputs)
      }

      case actions.RECEIVE_UPLOADERS: {
        return state
          // .set('uploaders', action.uploaders)
          // .setIn(
          //   ['currentExpedition', 'uploaders'],
          //   action.uploaders.toList()
          // )
      }

      default:
        return state
    }
  // } catch (err) {
  //   console.log('error in reducer:', err)
  // }
}

export default expeditionReducer
