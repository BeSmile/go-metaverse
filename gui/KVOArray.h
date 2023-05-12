#import <Cocoa/Cocoa.h>

@interface KVOArray : NSObject

// Convenience accessor
- (NSArray *)currentData;

// For KVC compliance, publicly declared for readability
- (void)insertObject:(id)object inDataAtIndex:(NSUInteger)index;
- (void)removeObjectFromDataAtIndex:(NSUInteger)index;
- (id)objectInDataAtIndex:(NSUInteger)index;
- (NSArray *)dataAtIndexes:(NSIndexSet *)indexes;
- (NSUInteger)count;

@end