interface DeleteMessageReactionRequest {
  roomId: string
  messageId: string
}

export async function deleteMessageReaction({ roomId, messageId }: DeleteMessageReactionRequest) {
  await fetch(`${import.meta.env.VITE_APP_API_URL}/rooms/${roomId}/messages/${messageId}/react`, {
    method: 'DELETE'
  })
}