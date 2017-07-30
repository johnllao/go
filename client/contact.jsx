import React from "react"
import ReactDOM from "react-dom"

export default class Contact extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            id   : '',
            name : '',
            email: ''
        }

        this.handleChange = this.handleChange.bind(this)
        this.handleAdd = this.handleAdd.bind(this);
    }

    handleChange(event) {

        const value = event.target.value

        if (event.target.name == "id") {
            this.setState((prevState, props) => {
                return { id: value }
            });
            return
        }

        if (event.target.name == "name") {
            this.setState((prevState, props) => {
                return { name: value }
            });
            return
        }

        if (event.target.name == "email") {
            this.setState((prevState, props) => {
                return {email: value }
            });
            return
        }
    }

    handleAdd(event) {
        this.props.addContacts({
            id: parseInt(this.state.id),
            name: this.state.name,
            email: this.state.email
        })
    }

    render() {
        return (
            <form className="form-horizontal">
                <div className="form-group">
                    <label className="col-sm-1 control-label">Name</label>
                    <div className="col-sm-1">
                        <input type="text" className="form-control" name="id" value={this.state.id} onChange={this.handleChange} />
                    </div>
                    <div className="col-sm-10">
                        <input type="text" className="form-control" name="name" value={this.state.name} onChange={this.handleChange} />
                    </div>
                </div>
                <div className="form-group">
                    <label className="col-sm-1 control-label">Email</label>
                    <div className="col-sm-11">
                        <input type="text" className="form-control" name="email" value={this.state.email} onChange={this.handleChange} />
                    </div>
                </div>
                <div className="btn-group pull-right" role="group" aria-label="..." style={ { paddingTop: '1em', paddingBottom: '1em' } }>
                    <button type="button" className="btn btn-default" name="add" onClick={this.handleAdd}>
                        <span className="glyphicon glyphicon-plus"></span>
                    </button>
                    <button type="button" className="btn btn-default" name="save">
                        <span className="glyphicon glyphicon-floppy-disk"></span>
                    </button>
                </div>
            </form>
        )
    }
}