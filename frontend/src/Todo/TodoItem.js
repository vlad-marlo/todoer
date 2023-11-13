import "./Todo.css";
import React from "react";

const TodoItem = ({todo}) => {
    return (
        <li className={`todo ${todo.status}`}>
            <h3>{todo.id}</h3>
            <p>{todo.value}</p>
            <h4>Status</h4>
            <p>{todo.status}</p>
            <button onClick={() => {
                fetch(`/api/v1/tasks/${todo.id}/status?status=deleted`, {
                    method: "POST"
                }).catch((e) => {
                    alert(e)
                }).then(r => r.json())
                    .then((todo) => {
                        console.log(todo)
                    })
            }}>DELETE</button>
        </li>
    )
}

export default TodoItem;