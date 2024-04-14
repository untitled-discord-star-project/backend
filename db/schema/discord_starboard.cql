CREATE KEYSPACE IF NOT EXISTS discord WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'} AND durable_writes = true;
USE discord;

-- excludes unnecessary and optional fields
CREATE TABLE IF NOT EXISTS discord_starboard (
    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#Message
    message_id text PRIMARY KEY,
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

    -- https://pkg.go.dev/github.com/bwmarrin/discordgo#UserS
    user_username text,
    user_discriminator text,
    user_global_name text,
);

INSERT INTO discord_starboard (message_id, message_content, message_author_id, message_author_username, message_attachment_id, message_attachment_url, message_attachment_proxy_url, message_attachment_filename, message_attachment_width, message_attachment_height, message_attachment_size, message_attachment_ephemeral, message_reaction_count, message_reaction_emoji, emoji_id, emoji_roles, emoji_require_colons, emoji_managed, emoji_animated, emoji_available, sticker_id, sticker_name, sticker_format, user_username, user_discriminator, user_global_name) VALUES ('123', 'hello', '123', 'user', '123', 'url', 'proxy', 'filename', 100, 100, 100, false, 1, 'emoji', 'emoji_id', ['role'], false, false, false, false, 'sticker_id', 'sticker_name', 1, 'username', 'discriminator', 'global_name');

