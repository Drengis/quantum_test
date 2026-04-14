export interface User {
  id: string;
  tg_id: string;
  username?: string;
  first_name: string;
  last_name?: string;
  lang_code?: string;
  invited_by?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface CreateUserRequest {
  tg_id: string;
  username?: string;
  first_name: string;
  last_name?: string;
  lang_code?: string;
  invited_by?: string;
}

export type TelegramUser = {
  tg_id?: number;
  username?: string;
  first_name?: string;
  last_name?: string;
  language_code?: string;
};