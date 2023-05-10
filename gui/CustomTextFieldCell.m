#import <CustomTextFieldCell.h>

@implementation CustomTextFieldCell

- (NSRect)titleRectForBounds:(NSRect)theRect {
    NSRect titleFrame = [super titleRectForBounds:theRect];

    CGFloat fontHeight = [[[NSLayoutManager alloc] init] defaultLineHeightForFont:self.font];
    if (fontHeight < titleFrame.size.height) {
//         NSLog(@"dfg%f", fontHeight);
        titleFrame.origin.y = theRect.origin.y + (theRect.size.height - fontHeight) * 0.5f;
        titleFrame.size.height = fontHeight;
    }
    return titleFrame;

}

- (void) drawInteriorWithFrame:(NSRect)cFrame inView:(NSView*)cView {
  [super drawInteriorWithFrame:[self titleRectForBounds:cFrame] inView:cView];
}

@end
