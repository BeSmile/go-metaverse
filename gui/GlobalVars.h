#import <Cocoa/Cocoa.h>
#import <ViewController.h>

@interface GlobalVars : NSObject
{
    NSScrollView *_scrollView;
    NSTableView  *_tableView;
    ViewController  *_viewController;
}

+ (GlobalVars *)sharedInstance;

@property(strong, nonatomic, readwrite) NSScrollView *scrollView;
@property(strong, nonatomic, readwrite) NSTableView *tableView;
@property(strong, nonatomic, readwrite) ViewController *viewController;

@end
