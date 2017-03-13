
import { connect } from 'react-redux'
import ExpeditionPage from '../../../components/AdminPage/ExpeditionPage/ExpeditionPage'
import * as actions from '../../../actions'

const mapStateToProps = (state, ownProps) => {

  const projects = state.expeditions.get('projects')
  const expeditions = state.expeditions.get('expeditions')
  const user = state.expeditions.get('user')

  return {
    ...ownProps,
    expeditions,
    projects,
    user
  }
}

const mapDispatchToProps = (dispatch, ownProps, state) => {
  return {
    setExpeditionProperty (key, value) {
      return dispatch(actions.setExpeditionProperty(key, value))
    },
    requestSignOut () {
      return dispatch(actions.requestSignOut())
    }
    // updateExpedition (expedition) {
    //   return dispatch(actions.updateExpedition(expedition))
    // },
  }
}

const ExpeditionPageContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(ExpeditionPage)

export default ExpeditionPageContainer
