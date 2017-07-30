import React       from "react"
import ReactDOM    from "react-dom"
import Axios       from "axios"
import Contact     from "./contact.jsx"
import ContactList from "./contactlist.jsx"

class App extends React.Component {

    constructor(props) {
        super(props)

        this.state = { contacts : [] }

        this.addContacts = this.addContacts.bind(this)

        Axios.get('/api/contacts')
             .then(response => {
                 this.setState((prevState, props) => {
                    return { contacts: response.data } 
                });
             });
    }

    addContacts(c) {
        this.setState((prevState, props) => {
            return { contacts: this.state.contacts.push(c) }
        });
    }

    render() {
        
        return (
            <div>
                <Contact addContacts={ this.addContacts }/>
                <ContactList contacts={ this.state.contacts } />
            </div>
        )
    }
}

ReactDOM.render(<App />, document.querySelector("#container"))