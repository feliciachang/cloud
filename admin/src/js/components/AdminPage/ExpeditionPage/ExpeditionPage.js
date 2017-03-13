import React from 'react'
import { Link } from 'react-router'
import Navigation from '../../Navigation/Navigation'
import BreadCrumbs from './../BreadCrumbs'


class ExpeditionPage extends React.Component {
  constructor (props) {
    super(props)
    this.state = {

    }
  }

  render () {

    const {
      children,
      currentExpedition,
      params,
      errors,
      location,
      expeditions,
      projects,
      user,
      requestSignOut
    } = this.props

    return (
      <div id="admin-page" className="page">
        <Navigation 
          { ...params }
          expeditions={ expeditions }
          projects={ projects }
          requestSignOut={ requestSignOut }
          user={ user }
        />
        <div className="page-content">
          {/*
            <BreadCrumbs 
              { ...location }
              breadcrumbs={ breadcrumbs }
            />
          */}
          { children }
        </div>
      </div>
    )
  }
}

export default ExpeditionPage