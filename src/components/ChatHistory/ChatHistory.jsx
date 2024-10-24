import React,{Component} from "react";
import './ChatHistory.scss';
import Message from '../message'
class ChatHistory extends Component{
    render(){
        const messages = this.props.chatHistory.map(msg=><Message key={msg.timeStamp} message={msg.data}/>);
        console.log(messages)
        return (
            <div className="ChatHistory">
                <h2>Chat History</h2>
                {messages}
            </div>
        );
    };
}

export default ChatHistory;