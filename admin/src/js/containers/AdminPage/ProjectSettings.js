
import { connect } from 'react-redux'
import ProjectSettings from '../../components/AdminPage/ProjectSettings'
import * as actions from '../../actions'

const mapStateToProps = (state, ownProps) => {
  const currentProject = state.expeditions.get('currentProject')
  const expeditions = state.expeditions.get('expeditions')
    .filter((e) => {
      return currentProject.get('expeditions').includes(e.get('id'))
    })
  const name = currentProject.get('name')
  const description = currentProject.get('description')

  return {
    ...ownProps,
    currentProject,
    expeditions,
    name,
    description
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
