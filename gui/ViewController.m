#import <ViewController.h>
#import <GlobalVars.h>
#import <MessageWrapper.h>
#import <CustomTextFieldCell.h>

@implementation ViewController{
    NSTableView *_tableView;
    NSMutableArray *_dataArray;
}
#define kImageColumnIdentifier @"ImageColumn"
#define kTextColumnIdentifier @"TextColumn"


-(void) viewDidLoad {
     [super viewDidLoad];
    _dataArray = [NSMutableArray array];

    GlobalVars *globals = [GlobalVars sharedInstance];

    NSScrollView *scrollView = globals.scrollView;
    // 设置滚动样式
//     scrollView.scrollerStyle = NSScrollerKnobStyleDefault;
    scrollView.frame = self.view.bounds;

    _tableView = [[NSTableView alloc]initWithFrame:self.view.bounds];
    [_tableView setBackgroundColor:[NSColor clearColor]];
//     [_tableView setUsesAlternatingRowBackgroundColors:YES];
    [_tableView setGridColor:[NSColor clearColor]];

//     _tableView.style = NSTableViewStylePlain;

//     [_tableView setContentInset:NSMakeSize(0, 0)];
    [[_tableView enclosingScrollView] setDrawsBackground:NO];
    [[_tableView enclosingScrollView] setBorderType:NSNoBorder];

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
    [_tableView reloadData];

    self.view = scrollView;

    [scrollView setDocumentView:_tableView];
}

-(NSInteger)numberOfRowsInTableView:(NSTableView *)tableView{
    return _dataArray.count;

}

-(void)loadView {
	self.view = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 300, 360)];
}

-(id)tableView:(NSTableView *)tableView viewForTableColumn:(NSTableColumn *)tableColumn row:(NSInteger) row {
	NodeWrapper *nw = _dataArray[row];
	Node *n = nw->node;

    // 获取表格列的标识符
    NSString *identifier = [tableColumn identifier];
    if ([kImageColumnIdentifier isEqualToString:identifier]) {
        NSView *imageView = [[NSView alloc] initWithFrame:NSMakeRect(0, 0, 36, 64)];
        // 如果是第一列，创建并返回一个 NSImageView
        NSImageView *avatarImageView = [[NSImageView alloc] initWithFrame:NSMakeRect(0, (imageView.frame.size.height - 36) / 2, 36, 36)];
        avatarImageView.image = [[NSImage alloc]initWithContentsOfURL:[NSURL URLWithString:n-> AVATAR]];
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

        // the value for y: in the frame should be row.height / 2 - textField.height / 2
//         VerticallyCenteredTextField *field = [[VerticallyCenteredTextField alloc] initWithFrame:CGRectMake(0, 0, 246, 60)];
//
//         NSTextFieldCell *cell = [field cell];
// //         [cell setVerticalAlignment:NSCellImageVertCenter];
//
//         [field setStringValue:n->TEXT];
//         [field setStringValue:n->TEXT];
//         [field setLineBreakMode:NSLineBreakByWordWrapping];
//
//         [field setAlignment:NSTextAlignmentCenter];
//
// //         [field setBordered:false];
//         NSFont *font = [NSFont systemFontOfSize:12.0]; // 字体大小为 14
//         [field setFont:font]; // 设置字体
// //         [field setIsEditable:false];
//         field.backgroundColor = [NSColor clearColor];

		NSTextField *textField = [[NSTextField alloc] initWithFrame:CGRectMake(0, 0, 246, 60)];
        NSFont *font = [NSFont systemFontOfSize:12.0]; // 字体大小为 14
        [textField setFont:font]; // 设置字体
        textField.backgroundColor = [NSColor clearColor];

        CustomTextFieldCell *cell = [[CustomTextFieldCell alloc] init];
//         [cell setWraps:YES];
//         [cell setScrollable:YES];
        [cell setStringValue:n->TEXT];
        [cell setTextColor:[NSColor whiteColor]]; // 设置文本颜色

//         [cell setLineBreakMode:NSLineBreakByWordWrapping];
        [textField setCell:cell];

        [cellView addSubview: textField];
        return cellView;
    } else {
        return nil;
    }
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


- (void) myCustomMethodWithString:(NSString *) avatar nn:(NSString *) nn text:(NSString *) text type:(NSString *) type {
    Node *n = (Node *) malloc(sizeof(Node));
	if(_dataArray.count > 5) {
		[_dataArray removeObjectAtIndex:0];
	}
    n->NN = nn;
    n->TEXT = text;
    n->TYPE = type;
    n->AVATAR = avatar;
    NodeWrapper *nw = [[NodeWrapper alloc] initWithNode:n];
    [_dataArray addObject:nw];

   	[_tableView reloadData];
}
@end
    
