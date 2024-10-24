import React, { Component } from "react";
import './chatinput.scss'

class ChatInput extends Component {
    render() {
        return(
            <div className="Input">
                <input onKeyDown={this.props.send} placeholder="Write Your Message...."></input>
            </div>
        );
    };
}

export default ChatInput;