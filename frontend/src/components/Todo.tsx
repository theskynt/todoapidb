import EditIcon from "../assets/edit.svg";
import DeleteIcon from "../assets/delete.svg";
import { useState } from "react";

interface TodoProps {
  id: number;
  title?: string;
  status?: string;
  onEdit?: (id: number, title: string) => void;
  onDelete?: (id: number) => void;
  onStatus?: (id: number, status: string) => void;
}

interface ButtonProps {
  onClick?: () => void;
}

export const EditButton: React.FC<ButtonProps> = (props) => {
  return (
    <button onClick={props.onClick}>
      <img src={EditIcon} className=" w-7 h-7" />
    </button>
  );
};

export const DeleteButton: React.FC<ButtonProps> = (props) => {
  return (
    <button onClick={props.onClick}>
      <img src={DeleteIcon} className=" w-7 h-7" />
    </button>
  );
};

const Todo: React.FC<TodoProps> = (props) => {
  const [isEditing, setIsEditing] = useState(false);
  const [title, setTitle] = useState(props.title);

  const save = () => {
    if (!title) return;
    props.onEdit?.(props.id, title);
    setIsEditing(false);
  };

  return (
    <section className=" w-full flex h-[66px] p-2 justify-between items-center">
      <label className={`cursor-pointer label gap-4 ${isEditing && "hidden"}`}>
        <input
          type="checkbox"
          defaultChecked={props.status == "done"}
          className="checkbox "
          onChange={(edit) => {
            props.onStatus?.(props.id, edit.target.checked ? "done" : "todo");
          }}
        />

        <span>{props.title}</span>
      </label>

      <input
        type="text"
        placeholder="Type here"
        className={`input input-bordered ${!isEditing && "hidden"} w-1/2`}
        value={title}
        onChange={(e) => setTitle(e.target.value)}
      />

      <section className={`flex items-center gap-4 ${isEditing && "hidden"}`}>
        <EditButton onClick={() => setIsEditing(true)} />
        <DeleteButton onClick={() => props.onDelete?.(props.id)} />
      </section>

      <section className={`flex items-center gap-4 ${!isEditing && "hidden"}`}>
        <button className="btn btn-md btn-success btn-outline" onClick={save}>
          Save
        </button>
        <button className="btn btn-md btn-error btn-outline" onClick={() => setIsEditing(false)}>
          Cancel
        </button>
      </section>
    </section>
  );
};

export default Todo;
