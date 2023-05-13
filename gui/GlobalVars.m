#import "GlobalVars.h"
#import <GlobalVars.h>

@implementation GlobalVars

@synthesize viewController = _viewController;

+ (GlobalVars *)sharedInstance {
    static dispatch_once_t onceToken;
    static GlobalVars *instance = nil;
    dispatch_once(&onceToken, ^{
        instance = [[GlobalVars alloc] init];
    });
    return instance;
}

- (id)init {
    self = [super init];
    if (self) {
        _viewController = [[ViewController alloc] init];
    }
    return self;
}

@end
