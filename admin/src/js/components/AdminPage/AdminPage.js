import React, {PropTypes, Children} from 'react'
import { Link } from 'react-router'
import ModalContainer from '../../containers/AdminPage/Modal'

class AdminPage extends React.Component {

  render () {

    const { 
      params,
      errors,
      location,
      expeditions,
      projects,
      breadcrumbs,
      user,
      requestSignOut
    } = this.props

    const children = Children.map(this.props.children,
      (child) => React.cloneElement(child, {
        errors
      })
    )    
    
    return (
      <div
        id="admin-page"
        className="page"
      >
        <ModalContainer/>
        { children }
      </div>
    )
  }
}

AdminPage.propTypes = {
}

export default AdminPage