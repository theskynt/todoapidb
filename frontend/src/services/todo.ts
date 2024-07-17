const all = async () => {
    try {
        const response = await fetch(import.meta.env.VITE_API_URL + "todos");
        const data = await response.json();
        return data ?? [];
    } catch (err) {
        console.log(err);
        return [];
    }
};

const add = async (title: string) => {
    const response = await fetch(import.meta.env.VITE_API_URL + "todos", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, status: "todo" }),
    });
    const data = await response.json();
    return data;
};

const del = async (id: number) => {
    await fetch(import.meta.env.VITE_API_URL + `todos/${id}`, {
        method: "DELETE",
    });
}

const updateStatus = async (id: number, status: string) => {
    await fetch(import.meta.env.VITE_API_URL + `todos/${id}/actions/status`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ status }),
    });
}

const updateTitle = async (id: number, title: string) => {
    await fetch(import.meta.env.VITE_API_URL + `todos/${id}/actions/title`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ title }),
    });
}

export default { all, add, del, updateStatus, updateTitle };
