
import React from 'react'
import { Link } from 'react-router'

class ProjectSettings extends React.Component {
  
  render () {

    const {
      currentProject,
      expeditions,
      name,
      description,
      promptModalNewExpedition
    } = this.props

    return (
      <div id="project-settings" className="project-settings">
        <div className="header">
          <h2>
            { currentProject.get('name') }
          </h2>
          <div
            className="button"
            onClick={ () => { console.log('not implemented yet') } }
          >
            Delete Project
          </div>
        </div>
        <div className="columns-container">
          <div className="column">
            <div className="expeditions">
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
              <div
                className="button"
                onClick={ promptModalNewExpedition }
              >
                Create new expedition
              </div>
            </div>
          </div>
          <div className="column">
            <div className="settings">
              <h3>
                Settings
              </h3>
              <div className="field">
                <p className="label">
                  Name:
                </p>
                <input 
                  type="text"
                  className={
                    '' +
                    (!!name && name.toLowerCase() !== 'project name' ? '' : ' default')
                  }
                  value={name}
                  onFocus={(e) => {
                    if (!name || name === 'Project Name') {
                      setProjectProperty(['name'], '')
                    }
                  }}
                  onChange={(e) => {
                    setProjectProperty(['name'], e.target.value)
                  }}
                />
                <p className="error"></p>
              </div>
              <div className="field">
                <p className="label">
                  URL:
                </p>
                <p>
                  { `https://${ name }/fieldkit.org`}
                </p>
              </div>
              <div className="field">
                <p className="label">
                  Description:
                </p>
                <input 
                  type="text"
                  className={
                    '' +
                    (!!description && description.toLowerCase() !== 'project description' ? '' : ' default')
                  }
                  value={description}
                  onFocus={(e) => {
                    if (!description || description === 'Project Description') {
                      setProjectProperty(['description'], '')
                    }
                  }}
                  onChange={(e) => {
                    setProjectProperty(['description'], e.target.value)
                  }}
                />
                <p className="error"></p>
              </div>
            </div>
            <div className="users">
              <h3>
                Users
              </h3>
              <p>
                Not implemented yet...
              </p>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

export default ProjectSettings