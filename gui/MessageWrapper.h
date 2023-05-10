typedef struct Node {
	// 昵称
    NSString* NN;
    // 头像
    NSString* AVATAR;
    // 消息
    NSString* TEXT;
    // 消息类型
    NSString* TYPE;
} Node;

@interface NodeWrapper : NSObject {
   @public

   Node *node;
}

- (id) initWithNode:(Node *) n;
@end