#import <KVOArray.h>

@interface KVOArray ()

@property (nonatomic, strong) NSMutableArray *data;

@end

@implementation KVOArray
//  We'll use automatic notifications for this example
+ (BOOL)automaticallyNotifiesObserversForKey:(NSString *)key
{
    if ([key isEqualToString:@"data"]) {
        return YES;
    }
    return [super automaticallyNotifiesObserversForKey:key];
}

- (id)init
{
    self = [super init];
    if (self) {
        // This is the ivar which provides storage
        _data = [NSMutableArray array];
    }
    return self;
}

//  Just a convenience method
- (NSArray *)currentData
{
    return [self dataAtIndexes:[NSIndexSet indexSetWithIndexesInRange:NSMakeRange(0, [self count])]];
}

//  These methods enable KVC compliance
- (void)insertObject:(id)object inDataAtIndex:(NSUInteger)index
{
    self.data[index] = object;
}

- (void)removeObjectFromDataAtIndex:(NSUInteger)index
{
    [self.data removeObjectAtIndex:index];
}

- (id)objectInDataAtIndex:(NSUInteger)index
{
    return self.data[index];
}

- (NSArray *)dataAtIndexes:(NSIndexSet *)indexes
{
    return [self.data objectsAtIndexes:indexes];
}

- (NSUInteger)count
{
    return [self.data count];
}

@end
