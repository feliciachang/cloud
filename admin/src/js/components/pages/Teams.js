// @flow weak

import React, { Component } from 'react'
import { Route, Link } from 'react-router-dom'
import ReactModal from 'react-modal';

import { MembersTable } from '../shared/MembersTable';
import { TeamForm } from '../forms/TeamForm';
import { MemberForm } from '../forms/MemberForm';
import { FKApiClient } from '../../api/api';
import '../../../css/teams.css'

import { FormContainer } from '../containers/FormContainer';

import type { APIProject, APIExpedition, APITeam, APINewTeam, APINewMember, APIMember } from '../../api/types';

type Props = {
  project: APIProject;
  expedition: APIExpedition;

  match: Object;
  location: Object;
  history: Object;
}

export class Teams extends Component {

  props: Props;
  state: {
    teams: APITeam[],
    members: { [teamId: number]: APIMember[] },
    memberDeletion: ?{
      contents: React$Element<*>;
      userId: number;
      teamId: number;
    }
  }

  constructor (props: Props) {
    super(props);
    this.state = {
      teams: [],
      members: {},
      memberDeletion: null
    }

    this.loadTeams();
  }

  async loadTeams() {
    const teamsRes = await FKApiClient.get().getTeamsBySlugs(this.props.project.slug, this.props.expedition.slug);
    if (teamsRes.type === 'ok' && teamsRes.payload) {
      const { teams } = teamsRes.payload;
      this.setState({ teams: teams });
      for (const team of teams) {
        await this.loadMembers(team.id);
      }
    }
  }

  async loadMembers(teamId: number) {
    const membersRes = await FKApiClient.get().getMembers(teamId);
    if (membersRes.type === 'ok' && membersRes.payload) {
      const { members } = this.state;
      members[teamId] = membersRes.payload.members;
      this.setState({members: members});
    }
  }

  async onTeamCreate(e: APINewTeam) {

    const { expedition, match } = this.props;

    const teamRes = await FKApiClient.get().createTeam(expedition.id, e);
    if (teamRes.type === 'ok') {
      await this.loadTeams();
      this.props.history.push(`${match.url}`);
    } else {
      return teamRes.errors;
    }
  }

  async onMemberAdd(teamId: number, e: APINewMember) {
    const { match } = this.props;
    const memberRes = await FKApiClient.get().addMember(teamId, e);
    if (memberRes.type === 'ok') {
      await this.loadTeams();
      this.props.history.push(`${match.url}`);
    } else {
      return memberRes.errors;
    }
  }

  startMemberDelete(teamId: number, userId: number) {
    const { teams, members } = this.state;
    let teamName = 'unknownTeam';
    let userName = 'unknownUser';

    if (members[teamId]) {
      const teamMembers = members[teamId];
      const teamMember = teamMembers.find((u: APIMember) => u.user_id === userId);
      const team = teams.find((t: APITeam) => t.id === teamId);

      if (team) {
        teamName = team.name;
      }
      if (teamMember) {
        // TODO: look up username
        userName = teamMember.user_id.toString(10);
      }
    }
    this.setState({
      memberDeletion: {
        contents: <span>Are you sure you want to remove <strong>{userName}</strong> from <strong>{teamName}</strong></span>,
        userId, teamId
      }
    })
  }

  async confirmMemberDelete() {
    const { memberDeletion } = this.state;

    if (memberDeletion) {
      const { teamId, userId } = memberDeletion;

      const memberRes = await FKApiClient.get().deleteMember(teamId, userId);
      if (memberRes.type === 'ok') {
        await this.loadMembers(teamId);
        this.setState({ memberDeletion: null })
      } else {
        // TODO: show errors somewhere
      }
    }
  }

  cancelMemberDelete() {
    this.setState({ memberDeletion: null });
  }

  render() {
    const { match } = this.props;
    const { members, teams, memberDeletion } = this.state;

    return (
      <div className="teams">
        <Route path={`${match.url}/new-team`} render={() =>
          <ReactModal isOpen={true} contentLabel="Create New Team">
            <h1>Create a new team</h1>
            <TeamForm
              onCancel={() => this.props.history.push(`${match.url}`)}
              onSave={this.onTeamCreate.bind(this)} />
          </ReactModal> } />

        <Route path={`${match.url}/:teamId/add-member`} render={props =>
          <ReactModal isOpen={true} contentLabel="Add Members">
            <h1>Add Members</h1>
            <MemberForm
              teamId={props.match.params.teamId}
              members={this.state.members[props.match.params.teamId]}
              onCancel={() => this.props.history.push(`${match.url}`)}
              onSave={this.onMemberAdd.bind(this)} />
          </ReactModal> } />

        { memberDeletion &&
          <ReactModal isOpen={true} contentLabel="Remove member confirmation">
            <h1>Confirm removal</h1>
            <FormContainer
              onSave={this.confirmMemberDelete.bind(this)}
              onCancel={this.cancelMemberDelete.bind(this)}
              saveText="Confirm"
            >
              <div>{memberDeletion.contents}</div>
            </FormContainer>
          </ReactModal> }

        <h1>Teams</h1>

        <div className="accordion">
        { teams.map((team, i) =>
          <div key={i} className="accordion-row expanded">
              <div className="accordion-row-header">
                <button className="secondary">Edit</button>
                <h4>{team.name}</h4>
                <p>{team.description}</p>
              </div>
              <div>
                { <MembersTable
                    teamId={team.id}
                    members={members[team.id]}
                    onDelete={this.onMemberDelete.bind(this)} /> }
                { !members[team.id] &&
                  <span className="empty">No members</span> }
                <Link className="button secondary" to={`${match.url}/${team.id}/add-member`}>Add Member</Link>                
              </div>
          </div> ) }
        { teams.length === 0 &&
          <span className="empty">No teams</span> }
        </div>
        <Link className="button" to={`${match.url}/new-team`}>Create New Team</Link>
      </div>
    )
  }
}
