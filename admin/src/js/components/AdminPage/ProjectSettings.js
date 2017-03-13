
import React from 'react'
import { Link } from 'react-router'

class ProjectSettings extends React.Component {
  
  render () {

    const {
      currentProject,
      expeditions
    } = this.props

    return (
      <div id="project-settings" className="project-settings">
        <div className="header">
          <h2>
            { currentProject.get('name') }
          </h2>
          <div className="button">
            Delete Project
          </div>
        </div>
        <div className="columns-container">
          <div className="column">
            <h3>
              Expeditions
            </h3>
            {
              !!expeditions &&
              expeditions.map((e, i) => {
                return (
                  <Link
                    to={ '/admin/' + currentProject.get('id') + '/' + e.get('id') }
                  >
                    <div
                      className="expedition"
                      key={ i }
                    >
                      <h4>
                        { e.get('name') }
                      </h4>
                    </div>
                  </Link>
                )
              })
            }
          </div>
          <div className="column">
            <h3>
              Settings
            </h3>
          </div>
        </div>
      </div>
    )
  }
}

export default ProjectSettings