import { test, expect } from '@playwright/test'

test('should response todo when post api/v1/todos', async ({
    request,
}) => {
    const data = {
        title: "Test Todo",
        status: "Active"
    }
    const reps = await request.post('api/v1/todos', { data })

    expect(reps.ok()).toBeTruthy()
    expect(await reps.json()).toEqual(
        expect.objectContaining({
            id: expect.any(Number),
            title: data.title,
            status: data.status
        })
    )
})

test('should response todo when put api/v1/todos', async ({
    request,
}) => {
    const postData = {
        title: "Test Todo",
        status: "Active"
    }
    const repsPost = await request.post('api/v1/todos', { data: postData })
    const postJson = await repsPost.json()

    const putData = {
        title: "Test Update Todo",
        status: "Inactive"
    }
    const reps = await request.put('api/v1/todos/' + postJson.id, { data: putData })

    expect(reps.ok()).toBeTruthy()
    expect(await reps.json()).toEqual(
        expect.objectContaining({
            id: expect.any(Number),
            title: putData.title,
            status: putData.status
        })
    )
})
