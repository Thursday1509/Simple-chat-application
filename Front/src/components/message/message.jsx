import React, { Component } from "react";
import './message.scss'

class Message extends Component{
    constructor(props){
        super(props);
        let temp = JSON.parse(this.props.Message);
        this.state={
            message:temp
        }
    }
    render(){
        return(
            <div className="Mesage">
                {this.state.message.body}
            </div>
        )
    }
}

export default Message;
