// Code generated by goctl. DO NOT EDIT.
package types

type ChangelogResponse struct {
	Changelog string `json:"changelog"`
}

type Chat struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UpdatedAt int    `json:"updated_at"`
	CreatedAt int    `json:"created_at"`
}

type ChatEntity struct {
	Id        string                   `json:"id"`
	Title     string                   `json:"title"`
	Models    []string                 `json:"models"`
	Options   map[string]interface{}   `json:"options,optional"`
	Messages  []map[string]interface{} `json:"messages,optional"`
	History   map[string]interface{}   `json:"history,optional"`
	Tags      []map[string]interface{} `json:"tags,optional"`
	Timestamp int64                    `json:"timestamp"`
}

type ChatMessage struct {
	Chats Chat `json:"chats"`
}

type ChatRespone struct {
	Text string `json:"text"`
}

type ConfigResponse struct {
	Status                   bool                       `json:"status"`
	Name                     string                     `json:"name"`
	Version                  string                     `json:"version"`
	DefaultLocale            string                     `json:"default_locale"`
	Images                   bool                       `json:"images"`
	DefaultModels            interface{}                `json:"default_models"`
	DefaultPromptSuggestions []DefaultPromptSuggestions `json:"default_prompt_suggestions"`
	TrustedHeaderAuth        bool                       `json:"trusted_header_auth"`
}

type CreateRequest struct {
	UserId   string `json:"userId"`
	Title    string `json:"title"`
	Chat     string `json:"chat"`
	ShareId  string `json:"shareId"`
	Archived int64  `json:"archived"`
}

type CreateResponse struct {
	Id int64 `json:"id"`
}

type DefaultModels struct {
	Models string `json:"models"`
}

type DefaultPromptSuggestions struct {
	Title   []string `json:"title"`
	Content string   `json:"content"`
}

type DetailRequest struct {
	Id int64 `json:"id"`
}

type DetailResponse struct {
	Id       int64  `json:"id"`
	UserId   string `json:"userId"`
	Title    string `json:"title"`
	Chat     string `json:"chat"`
	ShareId  string `json:"shareId"`
	Archived int64  `json:"archived"`
}

type Details struct {
	Format            string      `json:"format"`
	Family            string      `json:"family"`
	Families          interface{} `json:"families"`
	ParameterSize     string      `json:"parameter_size"`
	QuantizationLevel string      `json:"quantization_level"`
}

type ListRequest struct {
}

type ListResponse struct {
	Data []DetailResponse `json:"data"`
}

type MessagesEntity struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ModelDetail struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Model      string  `json:"model"`
	ModifiedAt string  `json:"modified_at"`
	Size       int64   `json:"size"`
	Digest     string  `json:"digest"`
	Details    Details `json:"details"`
}

type ModelReponse struct {
	Models []ModelDetail `json:"models"`
}

type NewChatEntity struct {
	Model    string           `json:"model"`
	Messages []MessagesEntity `json:"messages"`
	Options  OptionsEntity    `json:"options"`
}

type NewChatRequest struct {
	Chat ChatEntity `json:"chat"`
}

type OptionsEntity struct {
}

type Prompt struct {
	Id      int64  `json:"id"`
	Command string `json:"command"`
	UserId  string `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PromptResponse struct {
	Prompts []Prompt `json:"prompts"`
}

type RemoveRequest struct {
	Id int64 `json:"id"`
}

type RemoveResponse struct {
}

type UpdateChat struct {
	Messages []UpdateMessages `json:"messages"`
	History  UpdateHistory    `json:"history"`
}

type UpdateChatRequest struct {
	Chat UpdateChat `json:"chat"`
}

type UpdateHistory struct {
	Messages  map[string]interface{} `json:"messages"`
	CurrentId string                 `json:"currentId"`
}

type UpdateMessages struct {
	Id          string   `json:"id"`
	ChildrenIds []string `json:"childrenIds"`
	Role        string   `json:"role"`
	Content     string   `json:"content"`
	Timestamp   int64    `json:"timestamp"`
}

type UpdateRequest struct {
	Chat string `json:"chat"`
}

type UpdateResponse struct {
}

type VersionResponse struct {
	Version string `json:"version"`
}
