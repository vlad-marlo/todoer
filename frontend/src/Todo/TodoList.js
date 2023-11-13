import './Todo.css';
import React from 'react';
import TodoItem from "./TodoItem";

const TodoList = (items) => {
    return (
        <div className="todolist">
            <h2>list</h2>
            <ul>
                {items.items.map(todo => {
                    return <TodoItem todo={todo} key={todo.id}/>
                })}
            </ul>
        </div>
    )
}

export default TodoList;