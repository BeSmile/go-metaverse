#import <Cocoa/Cocoa.h>

@interface ViewController : NSViewController <NSTableViewDelegate,NSTableViewDataSource>

// 定义你需要的属性和方法
- (void) myCustomMethodWithString:(NSString *) avatar nn:(NSString *) nn text:(NSString *) text type:(NSString *) type;

@end
