#import <Cocoa/Cocoa.h>
#import <ViewController.h>

@interface GlobalVars : NSObject
{
    ViewController  *_viewController;
}

+ (GlobalVars *)sharedInstance;

@property(strong, nonatomic, readwrite) ViewController *viewController;

@end
