import WS_ADDRESS from "@/config";
import {useMessageStore} from "@/stores/modules/chatMessage";

const messageStore = useMessageStore()
export const useWebSocket = (userId, targetId) => {
    const ws = new WebSocket(WS_ADDRESS + "/sendUserMsg?userId=" + userId + "&targetId=" + targetId);

    const init = () => {
        ws.onopen = () => {
            console.log('WebSocket连接成功');
        };
        ws.onmessage = (e) => {
            console.log('收到消息', e.data);
            console.log('收到', JSON.parse(e.data).ownerId)
            messageStore.addMessage(JSON.parse(e.data).ownerId, JSON.parse(e.data))
        };
        ws.onclose = () => {
            console.log('WebSocket连接关闭');
        };
        ws.onerror = () => {
            console.log('WebSocket连接发生错误');
        };
    }
    init();
    return {
        ws,
    }
}