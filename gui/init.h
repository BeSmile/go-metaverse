#import <Cocoa/Cocoa.h>
#import <GlobalVars.h>
#import <ViewController.h>
#import <MessageWrapper.h>

int StartApp(void) {
    GlobalVars *globals = [GlobalVars sharedInstance];

    [NSAutoreleasePool new];
    [NSApplication sharedApplication];
    [NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
    id menubar = [[NSMenu new] autorelease];
    id appMenuItem = [[NSMenuItem new] autorelease];
    [menubar addItem:appMenuItem];
    [NSApp setMainMenu:menubar];
    id appMenu = [[NSMenu new] autorelease];
    id appName = [[NSProcessInfo processInfo] processName];
    id quitTitle = [@"Quit " stringByAppendingString:appName];
    id quitMenuItem = [[[NSMenuItem alloc] initWithTitle:quitTitle
        action:@selector(terminate:) keyEquivalent:@"q"]
          	autorelease];
    [appMenu addItem:quitMenuItem];
    [appMenuItem setSubmenu:appMenu];

    id window = [[[NSWindow alloc] initWithContentRect:NSMakeRect(0, 0, 300, 360)
        styleMask:NSTitledWindowMask backing:NSBackingStoreBuffered defer:NO]
            autorelease];
    [window cascadeTopLeftFromPoint:NSMakePoint(20,20)];
    [window setTitle:appName];
    //[window setOpaque:NO];
    //[window setLevel:NSFloatingWindowLevel];

	// 透明背景
    [window setBackgroundColor:[NSColor clearColor]];
    [window makeKeyAndOrderFront:nil];
    [NSApp activateIgnoringOtherApps:YES];

	NSView *contentView = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 300, 360)];
//	NSView *scrollContentView = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 300, 500)];

//	[scrollView setDocumentView:scrollContentView];
	// 透明背景
	[contentView setWantsLayer:YES];
	[contentView.layer setBackgroundColor:[NSColor clearColor].CGColor];

//	NSImageView *avatarImageView = [[NSImageView alloc] initWithFrame:NSMakeRect(0, 0, 36, 36)];
//	avatarImageView.image = [[NSImage alloc]initWithContentsOfURL:[NSURL URLWithString:@"https://himg.bdimg.com/sys/portrait/item/public.1.45bc355c.DyX87lFJ9AlPYb2-g7HO3Q.jpg"]];
//	avatarImageView.imageScaling =   NSImageScaleNone;
//	avatarImageView.wantsLayer = YES;
//	avatarImageView.layer.masksToBounds = YES;
//	avatarImageView.layer.cornerRadius = 18.f;
//  avatarImageView.layer.borderWidth = 3.0;
//  avatarImageView.layer.borderColor = [[NSColor redColor] CGColor];
//	[scrollContentView addSubview:avatarImageView];

//	NSTextView *textView = [[NSTextView alloc] initWithFrame:CGRectMake(20, 20, 200, 40)];
//	textView.string = @"Hello, world!";
//	textView.backgroundColor = [NSColor clearColor];
//	textView.textColor = [NSColor whiteColor];

//	[scrollContentView addSubview:textView];

//	NSTextView *textView1 = [[NSTextView alloc] initWithFrame:CGRectMake(20, 40, 200, 40)];
//	textView1.string = @"Hello, world!2";
//	textView1.backgroundColor = [NSColor clearColor];
//	textView1.textColor = [NSColor whiteColor];

//	[scrollContentView addSubview:textView1];

    // 创建表格
//    NSMutableArray *_dataArray = [NSMutableArray array];
//    for (int i = 0; i< 20; i++) {
//        [_dataArray addObject:[NSString stringWithFormat:@"%d行数据", i]];
//    }

//    NSTableView *tableView = globals.tableView;
//    tableView.translatesAutoresizingMaskIntoConstraints = NO;
//    tableView.backgroundColor = [NSColor redColor];
//    tableView.intercellSpacing = NSMakeSize(2, 2);
//    tableView.headerView = [[NSTableHeaderView alloc] initWithFrame:NSMakeRect(0, 0, 0, CGFLOAT_MIN)];
//
//    NSTableColumn *column = [[NSTableColumn alloc]initWithIdentifier:@"test"];
//    [tableView addTableColumn:column];
//    tableView.dataSource = _dataArray;
//    [scrollView setDocumentView:tableView];

    ViewController *vc = globals.viewController;
//    [self addChildViewController:viewController];
	[contentView addSubview:vc.view];

	// 背景透明
	//[scrollView setBackgroundColor:[NSColor clearColor]];
	//[scrollView setDrawsBackground:NO];
	//[scrollContentView setWantsLayer:YES];

	//scrollView.translatesAutoresizingMaskIntoConstraints = NO;
	//[scrollView.leadingAnchor constraintEqualToAnchor:contentView.leadingAnchor constant:10].active = YES;
	//[scrollView.trailingAnchor constraintEqualToAnchor:contentView.trailingAnchor constant:-10].active = YES;
	//[scrollView.topAnchor constraintEqualToAnchor:contentView.topAnchor constant:10].active = YES;
	//[scrollView.bottomAnchor constraintEqualToAnchor:contentView.bottomAnchor constant:-10].active = YES;
	if (![NSThread isMainThread]) {
	    dispatch_async(dispatch_get_main_queue(), ^{
	        // 在主线程上更新拖动区域
	        [window setContentView:contentView];
	    });
	} else {
	    [window setContentView:contentView];
	}
    [NSApp run];
    return 0;
}

void InitDataSource(const char* avatar, const char* nn, const char* text, const char* type) {
	NSString *avatarStr = [NSString stringWithUTF8String:avatar];
	NSString *nnStr =     [NSString stringWithUTF8String:nn];
	NSString *textStr =   [NSString stringWithUTF8String:text];
	NSString *typeStr =   [NSString stringWithUTF8String:type];

    GlobalVars *globals = [GlobalVars sharedInstance];
	ViewController *vc = globals.viewController;
	[vc myCustomMethodWithString:avatarStr nn:nnStr text:textStr type:typeStr];
}