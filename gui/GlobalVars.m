#import "GlobalVars.h"
#import <GlobalVars.h>

@implementation GlobalVars

@synthesize scrollView = _scrollView;
@synthesize tableView = _tableView;
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
        _scrollView = [[NSScrollView alloc] initWithFrame:NSMakeRect(0, 0, 300, 360)];
        // 背景透明
		[_scrollView setWantsLayer:YES];
		[_scrollView.layer setBackgroundColor:[NSColor clearColor].CGColor];
	    [_scrollView setDrawsBackground:NO];
	    [_scrollView setBackgroundColor:[NSColor clearColor]];
	    [_scrollView setHasVerticalScroller:YES];
	    [_scrollView setAutohidesScrollers:YES];

        _tableView = [[NSTableView alloc]  initWithFrame:NSZeroRect];
        _viewController = [[ViewController alloc] init];
    }
    return self;
}

@end
