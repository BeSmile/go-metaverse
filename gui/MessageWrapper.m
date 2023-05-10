#import <Cocoa/Cocoa.h>
#import <MessageWrapper.h>

@implementation NodeWrapper

- (id) initWithNode:(Node *) n {
  self = [super init];
  if(self) {
     node = n;
  }
  return self;
}

- (void) dealloc {
  free(node);
  [super dealloc];
}

@end