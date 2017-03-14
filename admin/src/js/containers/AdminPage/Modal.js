
import { connect } from 'react-redux'
import Modal from '../../components/AdminPage/Modal'
import * as actions from '../../actions'
import I from 'immutable'

const mapStateToProps = (state, ownProps) => {

  const type = state.expeditions.getIn(['modal', 'type'])
  let properties = I.Map()
  switch (type) {
    case 'new project': {
      properties = properties
        .set('name', state.expeditions.getIn(['currentProject', 'name']))
        .set('description', state.expeditions.getIn(['currentProject', 'description']))
    }
  }

  return {
    type,
    properties
  }
}

const mapDispatchToProps = (dispatch, ownProps) => {
  return {
    setProjectProperty (key, value) {
      return dispatch(actions.setProjectProperty(key, value))
    },
    closeModalAndCancel () {
      return dispatch(actions.closeModalAndCancel())
    },
    closeModalAndSave () {
      return dispatch(actions.closeModalAndSave())
    }
  }
}

const ModalContainer = connect(
  mapStateToProps,
  mapDispatchToProps
)(Modal)

export default ModalContainer
