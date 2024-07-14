import { test, expect } from "@playwright/test";

test("should response todo when post api/v1/todos", async ({ request }) => {
  const data = {
    title: "Test Todo",
    status: "Active",
  };
  const reps = await request.post("api/v1/todos", { data });
  const postJson = await reps.json();

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      id: expect.any(Number),
      title: data.title,
      status: data.status,
    })
  );

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);
  expect(deleteResponse.ok()).toBeTruthy();
});

test("should response todo when get api/v1/todos/:id", async ({ request }) => {
  const postData = {
    title: "Test Todo",
    status: "Active",
  };
  const repsPost = await request.post("api/v1/todos", { data: postData });
  const postJson = await repsPost.json();

  const reps = await request.get("api/v1/todos/" + postJson.id);

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      id: expect.any(Number),
      title: postData.title,
      status: postData.status,
    })
  );

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);
  expect(deleteResponse.ok()).toBeTruthy();
});

test("should response todo when put api/v1/todos/:id", async ({ request }) => {
  const postData = {
    title: "Test Todo",
    status: "Active",
  };
  const repsPost = await request.post("api/v1/todos", { data: postData });
  const postJson = await repsPost.json();

  const putData = {
    title: "Test Update Todo",
    status: "Inactive",
  };
  const reps = await request.put("api/v1/todos/" + postJson.id, {
    data: putData,
  });

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      id: expect.any(Number),
      title: putData.title,
      status: putData.status,
    })
  );

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);
  expect(deleteResponse.ok()).toBeTruthy();
});

test("should response todo when delete api/v1/todos/:id", async ({
  request,
}) => {
  const data = {
    title: "Test Todo",
    status: "Active",
  };
  const reps = await request.post("api/v1/todos", { data });
  const postJson = await reps.json();

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);

  expect(deleteResponse.ok()).toBeTruthy();
  const deleteResult = await deleteResponse.json();
  expect(deleteResult).toEqual("succes");
});

test("should response todo when PATCH api/v1/todos/:id/actions/status", async ({
  request,
}) => {
  const postData = {
    title: "Test Todo",
    status: "Active",
  };
  const repsPost = await request.post("api/v1/todos", { data: postData });
  const postJson = await repsPost.json();

  const patchStatus = {
    status: "inactive",
  };
  const reps = await request.patch(
    "api/v1/todos/" + postJson.id + "/actions/status",
    {
      data: patchStatus,
    }
  );

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      id: expect.any(Number),
      title: expect.any(String),
      status: patchStatus.status,
    })
  );

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);
  expect(deleteResponse.ok()).toBeTruthy();
});

test("should response todo when PATCH api/v1/todos/:id/actions/title", async ({
  request,
}) => {
  const postData = {
    title: "Test Todo",
    status: "Active",
  };
  const repsPost = await request.post("api/v1/todos", { data: postData });
  const postJson = await repsPost.json();

  const patchTitle = {
    title: "todo 2",
  };
  const reps = await request.patch(
    "api/v1/todos/" + postJson.id + "/actions/title",
    {
      data: patchTitle,
    }
  );

  expect(reps.ok()).toBeTruthy();
  expect(await reps.json()).toEqual(
    expect.objectContaining({
      id: expect.any(Number),
      title: patchTitle.title,
      status: expect.any(String),
    })
  );

  const deleteResponse = await request.delete(`/api/v1/todos/${postJson.id}`);
  expect(deleteResponse.ok()).toBeTruthy();
});
