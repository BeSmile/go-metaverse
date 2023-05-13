#import <ViewController.h>
#import <GlobalVars.h>
#import <MessageWrapper.h>
#import <CustomTextFieldCell.h>
#import <KVOArray.h>

@interface ViewController ()

@property (weak) IBOutlet NSTableView * tableView;
@property (nonatomic,strong) KVOArray *dataArray;
@property (strong) NSOperationQueue * queue;

@end

@implementation ViewController{
    NSTableView *_tableView;
//     KVOArray *_dataArray;
}
#define kImageColumnIdentifier @"ImageColumn"
#define kTextColumnIdentifier @"TextColumn"

static char MyObservationContext;

- (id)initWithNibName:(NSString *)nibNameOrNil bundle:(NSBundle *)nibBundleOrNil
{
    self = [super initWithNibName:nibNameOrNil bundle:nibBundleOrNil];
    if (self) {
        //  Instantiate a DataContainer and store it in our property
        _dataArray = [[KVOArray alloc] init];
        //  Add self as an observer. The context is used to verify that code from this class (and not its superclass) started observing.
        [_dataArray addObserver:self
                         forKeyPath:@"data"
                            options:(NSKeyValueObservingOptionOld | NSKeyValueObservingOptionNew)
                            context:&MyObservationContext];
    }

    return self;
}

-(void) viewDidLoad {
    [super viewDidLoad];
//    self.dataArray = [[KVOArray alloc] init];
//     [self loadData];
    self.queue = [[NSOperationQueue alloc] init];

//    [_dataArray addObserver:self
//         forKeyPath:@"dataArray"
//            options:(NSKeyValueObservingOptionOld | NSKeyValueObservingOptionNew)
//            context:&MyObservationContext];

    GlobalVars *globals = [GlobalVars sharedInstance];

    NSScrollView *scrollView = globals.scrollView;
    // 设置滚动样式
//     scrollView.scrollerStyle = NSScrollerKnobStyleDefault;
    scrollView.frame = self.view.bounds;
    _tableView = [[NSTableView alloc]initWithFrame:self.view.bounds];
    
    [_tableView setBackgroundColor:[NSColor clearColor]];
//     [_tableView setUsesAlternatingRowBackgroundColors:YES];
    [_tableView setGridColor:[NSColor clearColor]];
//    [_tableView setPreservesSelection:@NO];
    
    scrollView.backgroundColor = _tableView.backgroundColor;
//     _tableView.style = NSTableViewStylePlain;
    scrollView.drawsBackground = NO;

//     [_tableView setContentInset:NSMakeSize(0, 0)];
    [[_tableView enclosingScrollView] setDrawsBackground:NO];
    [[_tableView enclosingScrollView] setBorderType:NSNoBorder];
    [_tableView.enclosingScrollView.contentView setNeedsDisplay:YES];
    [_tableView.enclosingScrollView.layer setBackgroundColor:[NSColor clearColor].CGColor];
//    _tableView.enclosingScrollView.layer.backgroundColor = [NSColor clearColor].CGColor;


    NSTableColumn *avatarColumn = [[NSTableColumn alloc]initWithIdentifier:kImageColumnIdentifier];
	[avatarColumn setWidth: 36];
    [_tableView addTableColumn:avatarColumn];
    NSTableColumn *textColumn = [[NSTableColumn alloc]initWithIdentifier:kTextColumnIdentifier];
    [textColumn setWidth: 264];
    [_tableView addTableColumn:textColumn];

     // 隐藏表头
    NSTableHeaderView *headerView = [[NSTableHeaderView alloc] initWithFrame:NSZeroRect];
    _tableView.headerView = headerView;

    _tableView.delegate = self;
    _tableView.dataSource = self;
//     [_tableView reloadData];

    self.view = scrollView;
//     self.tableView = _tableView;

    [scrollView setDocumentView:_tableView];
}

#pragma mark - KVO

//- (void)observeValueForKeyPath:(NSString *)keyPath ofObject:(id)object change:(NSDictionary<NSKeyValueChangeKey,id> *)change context:(void *)context;{
//NSLog(@"监听字段:%s", keyPath);
//    if ([keyPath isEqualToString:@"dataArray"]) {
//        dispatch_async(dispatch_get_main_queue(), ^{
//            NSInteger index = self.dataArray.count - 1;
//            NSLog(@"%d", index);
//            if (index >= 0) {
//                [self.tableView beginUpdates];
//                [self.tableView insertRowsAtIndexes:[NSIndexSet indexSetWithIndex:index] withAnimation:NSTableViewAnimationSlideDown];
//                [self.tableView endUpdates];
//            }
//        });
//    }
//}

- (void)observeValueForKeyPath:(NSString *)keyPath ofObject:(id)object change:(NSDictionary *)change context:(void *)context
{
    //  Check if our class, rather than superclass or someone else, added as observer
    
    if (context == &MyObservationContext) {
        //  Check that the key path is what we want
        if ([keyPath isEqualToString:@"data"]) {
            //  Verify we're observing the correct object
            if (object == self.dataArray) {
//                NSLog(@"KVO for our container property, change dictionary is %@", change);
                dispatch_async(dispatch_get_main_queue(), ^{
                    [_tableView reloadData];
                    if (@available(macOS 10.14, *)) {
                        [_tableView performSelector:@selector(setNeedsDisplay:) withObject:@YES];
                    }
                    [_tableView setNeedsDisplay:YES];
                    [_tableView beginUpdates];
                    [_tableView endUpdates];
                });
            }
        }
    }
    else {
        //  Otherwise, call up to superclass implementation
        [super observeValueForKeyPath:keyPath ofObject:object change:change context:context];
    }
}


-(NSInteger)numberOfRowsInTableView:(NSTableView *)tableView{
    return self.dataArray.count;

}

-(void)loadView {
	self.view = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 300, 360)];
}

-(id)tableView:(NSTableView *)tableView viewForTableColumn:(NSTableColumn *)tableColumn row:(NSInteger) row {
// 	NSIndexSet *dataIndex = [NSIndexSet indexSetWithIndex: row];
	NodeWrapper *nw = [self.dataArray objectInDataAtIndex: row];
	Node *n = nw->node;

    // 获取表格列的标识符
    NSString *identifier = [tableColumn identifier];
    if ([kImageColumnIdentifier isEqualToString:identifier]) {
        NSString *avatar = n-> AVATAR;
       
        NSView *imageView = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 36, 60)];
        if( [avatar isEqualToString:@""]) {
            return imageView;
        }
        // 如果是第一列，创建并返回一个 NSImageView
        NSImageView *avatarImageView = [[NSImageView alloc] initWithFrame:NSMakeRect(0, (imageView.frame.size.height - 36) / 2, 36, 36)];
        avatarImageView.image = [[NSImage alloc]initWithContentsOfURL:[NSURL URLWithString:avatar]];
//         avatarImageView.imageScaling =   NSImageScaleNone;
        avatarImageView.wantsLayer = YES;
        avatarImageView.layer.masksToBounds = YES;
        avatarImageView.layer.cornerRadius = 18.f;
//         avatarImageView.layer.borderWidth = 1.0;
        [imageView addSubview: avatarImageView];

        [imageView setWantsLayer:YES];
        [imageView.layer setBackgroundColor:[NSColor clearColor].CGColor];

        return imageView;
    } else if ([kTextColumnIdentifier isEqualToString:identifier]) {
		NSTableCellView *cellView = [[NSTableCellView alloc] init];
		NSTextField *textField = [[NSTextField alloc] initWithFrame:CGRectMake(0, 0, 220, 60)];
        NSFont *font = [NSFont systemFontOfSize:12.0]; // 字体大小为 14
        [textField setFont:font]; // 设置字体
        [textField setTextColor:[NSColor clearColor]];
        textField.backgroundColor = [NSColor clearColor];

        CustomTextFieldCell *cell = [[CustomTextFieldCell alloc] init];
        [cell setStringValue:n->TEXT];
        NSColor *color = [self rainbowColorAtIndex:row]; // 生成第一个彩虹颜色

        [cell setTextColor:color]; // 设置文本颜色

        [textField setCell:cell];
        cellView.wantsLayer = YES;

        cellView.layer.backgroundColor = [NSColor clearColor].CGColor;

        [cellView addSubview: textField];
        return cellView;
    } else {
        return nil;
    }
}

- (void)tableView:(NSTableView *)tableView didAddRowView:(NSTableRowView *)rowView forRow:(NSInteger)row {
    [rowView setNeedsDisplay:YES];
}

- (CGFloat)tableView:(NSTableView *)tableView heightOfRow:(NSInteger)row {
    if (row == -1) {
        // 如果是表头行，设置其高度为 30
        return 30;
    } else {
        // 如果是数据行，设置其高度为 40
        return 60;
    }
}

- (void)loadData {
    NSOperationQueue *queue = [[NSOperationQueue alloc] init];

    [queue addOperationWithBlock:^{
        while (true) {
            // 模拟更新数据源过程，假设每秒更新一次
//             [NSThread sleepForTimeInterval:1];

            // 每秒更新一次数据源
//             NSString newData = [[NSDate date] descriptionWithLocale:[NSLocale currentLocale]];
//             [self.dataSource addObject:newData];

            // 使用主线程更新 table view
            [[NSOperationQueue mainQueue] addOperationWithBlock:^{
                
                
//                NSIndexSet *indexSetToInsert = [NSIndexSet indexSetWithIndex:self.dataArray.count - 1];
//                if(self.total >= 6) {
//
//                    NSIndexSet *indexSetToDelete = [NSIndexSet indexSetWithIndex:0];
//                    [_tableView beginUpdates];
//
//                    // NSTableViewAnimationEffectNone
//                    [_tableView removeRowsAtIndexes:indexSetToDelete withAnimation:NSTableViewAnimationEffectFade];
//                    [_tableView endUpdates]; // 结束动画代码块
//
//                }
//                [_tableView beginUpdates];
//                [_tableView insertRowsAtIndexes:indexSetToInsert withAnimation:NSTableViewAnimationSlideUp];// NSTableViewAnimationSlideUp  NSTableViewAnimationSlideRight
//                // 使用 beginUpdates 和 endUpdates 的方式更新 table view
//                [_tableView endUpdates];
                [_tableView reloadData];
//
//                 [_tableView beginUpdates];
////
// 				NSIndexSet *indexSet = [[NSIndexSet alloc] initWithIndexesInRange:NSMakeRange(0, _dataArray.count)];
// 				[_tableView reloadSections:indexSet];
////
//                 [_tableView endUpdates];
            }];
        }
    }];
}

- (void)dealloc
{
	[super dealloc];
    [self.dataArray removeObserver:self forKeyPath:@"data" context:&MyObservationContext];
}

- (NSColor *)rainbowColorAtIndex:(NSInteger)index {
//    CGFloat hue = (CGFloat)index / 6.0f;
//    return [NSColor colorWithHue:hue saturation:1.0f brightness:1.0f alpha:1.0f];
    CGFloat hue = (CGFloat)index / 6.0f;
    CGFloat saturation = 1.0f - (arc4random_uniform(30) / 100.0f);
    CGFloat brightness = 1.0f - (arc4random_uniform(30) / 100.0f);
    return [NSColor colorWithHue:hue saturation:saturation brightness:brightness alpha:1.0f];
}

- (void) myCustomMethodWithString:(NSString *) avatar nn:(NSString *) nn text:(NSString *) text type:(NSString *) type {
	@synchronized (self) { // 添加锁
       if (self.dataArray.count > 5) {
	       [self.dataArray removeObjectFromDataAtIndex:0];
	   }
    
	   Node *n = (Node *) malloc(sizeof(Node));
	   n->NN = nn;
	   n->TEXT = text;
	   n->TYPE = type;
	   n->AVATAR = avatar;
	   NodeWrapper *nw = [[NodeWrapper alloc] initWithNode:n];
//       [self.dataArray insertObject:[NSObject new] inDataAtIndex:0];

	   [self.dataArray insertObject:nw inDataAtIndex:self.dataArray.count];
//         NSIndexSet *indexSetToInsert = [NSIndexSet indexSetWithIndex:_dataArray.count - 1];
//         [_tableView reloadData];
// 		[_tableView beginUpdates];
//         [_tableView insertRowsAtIndexes:indexSetToInsert withAnimation:NSTableViewAnimationSlideUp];// NSTableViewAnimationSlideUp  NSTableViewAnimationSlideRight
//         [_tableView endUpdates]; // 结束动画代码块
	}
        // reloadData 方法放在动画代码块外
}

- (BOOL)tableView:(NSTableView *)tableView shouldSelectRow:(NSInteger)row {
    return NO;
}

@end
    
