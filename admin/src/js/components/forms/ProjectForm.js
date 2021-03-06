// @flow weak

import React, { Component } from 'react'

import { FormContainer } from '../containers/FormContainer';
import { errorsFor, slugify, fkHost } from '../../common/util';

import type { APIErrors, APINewProject } from '../../api/types';

type Props = {
name?: string,
slug?: string,
description?: string,

cancelText?: string;
saveText?: ?string;
onCancel?: () => void;
onSave: (p: APINewProject) => Promise<?APIErrors>;
}

export class ProjectForm extends Component {
    props: Props;
    state: {
    name: string,
    slug: string,
    description: string,
    slugHasChanged: boolean,
    saveDisabled: boolean,
    errors: ?APIErrors
    };

    constructor(props: Props) {
        super(props)
        const {name, slug, description} = props;

        this.state = {
            name: name || '',
            slug: slug || '',
            description: description || '',
            slugHasChanged: !!name && !!slug && slugify(name) != slug,
            saveDisabled: true,
            errors: null
        }
    }

    componentWillReceiveProps(nextProps: Props) {
        const {name, slug, description} = nextProps;

        this.setState({
            name: name || '',
            slug: slug || '',
            description: description || '',
            slugHasChanged: !!name && !!slug && slugify(name) != slug,
            saveDisabled: true,
            errors: null
        });
    }

    async save() {
        const errors = await this.props.onSave({
            name: this.state.name,
            slug: this.state.slug,
            description: this.state.description
        });
        if (errors) {
            this.setState({
                errors
            });
        }
    }

    handleNameChange(event) {
        const v = event.target.value;
        let stateUpdate = {
            name: v,
            saveDisabled: false
        };
        if (!this.state.slugHasChanged) {
            stateUpdate = {
                slug: slugify(v),
                ...stateUpdate
            } ;
        }
        this.setState(stateUpdate);
    }

    handleSlugChange(event) {
        const v = event.target.value;
        this.setState({
            slug: v,
            slugHasChanged: true,
            saveDisabled: false
        });
    }

    handleInputChange(event) {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        const name = target.name;
        this.setState({
            [name]: value,
            saveDisabled: false
        });
    }

    render() {
        const {saveDisabled} = this.state;
        return (
            <FormContainer onSave={ this.save.bind(this) } onCancel={ this.props.onCancel } saveText={ this.props.saveText } cancelText={ this.props.cancelText } saveDisabled={ saveDisabled }>
                <div className="form-group">
                    <label htmlFor="name">Name</label>
                    <input type="text" name="name" className='lg' value={ this.state.name } onChange={ this.handleNameChange.bind(this) } />
                    { errorsFor(this.state.errors, 'name') }
                </div>
                <div className="url-preview">
                    <p className="label">Your project will be available at:</p>
                    <p className="url">
                        <input type="text" name="slug" className='slug' value={ this.state.slug } onChange={ this.handleSlugChange.bind(this) } />
                        { /* TODO: replace with something that handles alternative domains */ }
                        <span className="domain">.{ fkHost() }/</span>
                    </p>
                    { errorsFor(this.state.errors, 'slug') }
                </div>
                <div className="form-group">
                    <label htmlFor="description">Description</label>
                    <input type="text" name="description" className='lg' value={ this.state.description } onChange={ this.handleInputChange.bind(this) } />
                    { errorsFor(this.state.errors, 'description') }
                </div>
            </FormContainer>
            );
    }
}
