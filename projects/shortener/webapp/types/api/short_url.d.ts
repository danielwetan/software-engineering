interface CreateShortUrlPayload {
  target: string;
}

interface ShortUrl {
  shortcode: string;
  user_id: int;
  target: string;
  created_at: string;
}
