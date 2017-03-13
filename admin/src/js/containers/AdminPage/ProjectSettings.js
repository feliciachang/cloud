
import { connect } from 'react-redux'
import ProjectSettings from '../../components/AdminPage/ProjectSettings'
import * as actions from '../../actions'

const mapStateToProps = (state, ownProps) => {
  const exp = state.expeditions
  // const projectID = expeditions.getIn(['currentProject', 'id'])
  // const expedition = expeditions.get('currentExpedition')

  const currentProject = exp.get('currentProject')
  const expeditions = exp.get('expeditions')
    .filter((e) => {
      return currentProject.get('expeditions').includes(e.get('id'))
    })

  return {
    ...ownProps,
    currentProject,
    expeditions
  }
}

const mapDispatchToProps = (dispatch, ownProps, state) => {
  return {
    setExpeditionProperty (keyPath, value) {
      return dispatch(actions.setExpeditionProperty(keyPath, value))
    }
  }
}

const ProjectSettingsContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ProjectSettings)

export default ProjectSettingsContainer
