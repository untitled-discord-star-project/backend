-- excludes unnecessary and optional fields
CREATE TABLE IF NOT EXISTS discord_starboard (
    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#Message
    message_id text,
    message_content text,
    message_author_id text,
    message_author_username text,
    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#MessageAttachment
    message_attachment_id text,
    message_attachment_url text,
    message_attachment_proxy_url text,
    message_attachment_filename text,
    message_attachment_width int,
    message_attachment_height int,
    message_attachment_size int,
    message_attachment_ephemeral boolean,

    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#MessageAttachment
    message_reaction_count int,
    message_reaction_emoji text,

    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#Emoji
    emoji_id text,
    emoji_roles list<text>,
    emoji_require_colons boolean,
    emoji_managed boolean,
    emoji_animated boolean,
    emoji_available boolean,

    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#StickerItem
    sticker_id text,
    sticker_name text,
    sticker_format int,

    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#User
    user_username text,
    user_discriminator text,
    user_global_name text,
    PRIMARY KEY ((message_id, user_global_name))
);