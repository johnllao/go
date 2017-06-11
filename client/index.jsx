import React from "react"
import ReactDOM from "react-dom"

class App extends React.Component {
    render() {
        return (
            <div>
                Main
            </div>
        )
    }
}

ReactDOM.render(<App />, document.querySelector("#container"))