import React, { Component } from 'react'
import Table from './Table'
import Form from './Form'

class App extends Component {
    state = {
        characters: [],
        data: [],
   }


    removeCharacter = (index) => {
        const {characters} = this.state
        this.setState({
            characters: characters.filter((character, i) => {
                return i !== index
            }),
        })
    }

    handleSubmit = (character) => {
        this.setState({ characters: [...this.state.characters,character]})
    }
    render() {
        const {characters} = this.state


      return(
        <div className="App">
          <h1>hello,React</h1>
          <div className="container">
                <Table characterData={characters} removeCharacter={this.removeCharacter}/>
            </div>
            <Form handleSubmit={this.handleSubmit} />
        </div>
        
      )
    }
  }


export default App