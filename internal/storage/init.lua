box.cfg {
    listen = 3301,
    memtx_memory = 512 * 1024 * 1024
}

box.once("schema", function()
    local space = box.schema.space.create("polls")
    space:format({
        {name = "id", type = "string"},
        {name = "creator_id", type = "string"},
        {name = "question", type = "string"},
        {name = "options", type = "array"},
        {name = "votes", type = "map"},
        {name = "is_closed", type = "boolean"},
    })

    space:create_index("primary", {
        type = "hash",
        parts = {"id"},
    })
end)
