import React from "react"
import ReactDOM from "react-dom"

export default class ContactList extends React.Component {

    constructor(props) {
        super(props)
    }

    render() {

        return (
            <table className="table table-hover">
                <thead>
                    <tr>
                        <th className="col-sm-1">ID</th>
                        <th className="col-sm-6">Name</th>
                        <th className="col-sm-5">Email</th>
                    </tr>
                </thead>
                <tbody>
                    {this.props.contacts.map(c => 
                        <tr key={ c.id }>
                            <td className="col-sm-1">{ c.id }</td>
                            <td className="col-sm-6">{ c.name }</td>
                            <td className="col-sm-5">{ c.email }</td>
                        </tr>
                    )}
                </tbody>
            </table>
        )
    }
}