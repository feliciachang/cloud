
import { connect } from 'react-redux'
import ProjectGrid from '../../components/AdminPage/ProjectGrid'
import * as actions from '../../actions'

const mapStateToProps = (state, ownProps) => {
  const expeditions = state.expeditions
  const projects = expeditions.get('projects')
  const user = expeditions.get('user')
  // const projectID = expeditions.getIn(['currentProject', 'id'])
  // const expedition = expeditions.get('currentExpedition')

  return {
    ...ownProps,
    projects,
    user
  }
}

const mapDispatchToProps = (dispatch, ownProps, state) => {
  return {
    setExpeditionProperty (keyPath, value) {
      return dispatch(actions.setExpeditionProperty(keyPath, value))
    },
    promptModalNewProject () {
      return dispatch(actions.promptModal('new project'))
    }
  }
}

const ProjectGridContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ProjectGrid)

export default ProjectGridContainer
