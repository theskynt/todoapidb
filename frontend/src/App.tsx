import { useEffect, useState } from "react";
import "./App.css";
import Todo from "./components/Todo";
import todoServices from "./services/todo";

interface Todo {
  id: number;
  title: string;
  status: string;
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [newTodo, setNewTodo] = useState<string>("");

  useEffect(() => {
    onLoadTodo();
  }, []);

  const onLoadTodo = async () => {
    const data = await todoServices.all();
    setTodos(data);
  };

  const onAddTodo = async () => {
    await todoServices.add(newTodo);
    setNewTodo("");
    onLoadTodo();
  };

  const onDeleteTodo = async (id: number) => {
    await todoServices.del(id);
    onLoadTodo();
  };

  const onUpdateStatus = async (id: number, status: string) => {
    await todoServices.updateStatus(id, status);
    onLoadTodo();
  };

  const onUpdateTitle = async (id: number, title: string) => {
    await todoServices.updateTitle(id, title);
    onLoadTodo();
  };

  return (
    <section className=" relative md:w-[650px] w-full h-full bg-white rounded-3xl flex flex-col overflow-hidden">
      <section className=" p-6 border-b">
        <h1> Todo List</h1>
      </section>

      <section className=" px-6 flex-1 flex-col overflow-scroll">
        {todos.map((todo) => (
          <Todo {...todo} key={todo.id} onDelete={onDeleteTodo} onStatus={onUpdateStatus} onEdit={onUpdateTitle} />
        ))}
      </section>

      <section className=" p-6 border-t">
        <input
          type="text"
          placeholder="Type here"
          className="input input-bordered w-full "
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
        />
        <button className="btn w-full mt-4" onClick={onAddTodo}>
          Add Todo
        </button>
      </section>
    </section>
  );
}

export default App;
