
import { connect } from 'react-redux'
import NewExpeditionPage from '../../../components/AdminPage/NewExpeditionPage/NewExpeditionPage'
import * as actions from '../../../actions'

const mapStateToProps = (state, ownProps) => {

  state = state.expeditions
  const projectID = state.getIn(['currentProject', 'id'])
  const expeditions = state.get('expeditions')
  const expedition = state.get('currentExpedition')

  return {
    ...ownProps,
    projectID,
    expeditions,
    expedition
  }
}

const mapDispatchToProps = (dispatch, ownProps, state) => {
  return {
    setExpeditionProperty (key, value) {
      return dispatch(actions.setExpeditionProperty(key, value))
    },
    setExpeditionPreset (type) {
      return dispatch(actions.setExpeditionPreset(type))
    },
    saveGeneralSettings (callback) {
      return dispatch(actions.saveGeneralSettings(callback))
    }
  }
}

const NewGeneralSettingsContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(NewExpeditionPage)

export default NewGeneralSettingsContainer