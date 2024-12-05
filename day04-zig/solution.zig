const std = @import("std");

fn readFileAndParse(comptime path: []const u8, allocator: std.mem.Allocator) !std.ArrayList(std.ArrayList(u8)) {
    const file = @embedFile(path);
    var it = std.mem.tokenize(u8, file, "\n");
    var grid = std.ArrayList(std.ArrayList(u8)).init(allocator);

    while (it.next()) |line| {
        var row = std.ArrayList(u8).init(allocator);
        try row.appendSlice(line);
        try grid.append(row);
        std.debug.print("Line: {s}\n", .{line});
    }

    return grid;
}

test "foo" {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const grid = try readFileAndParse("test-input.txt", allocator);

    defer {
        for (grid.items) |row| {
            row.deinit();
        }
        grid.deinit();
    }

    for (grid.items, 0..) |row, i| {
        for (row.items, 0..) |cell, j| {
            std.debug.print("grid[{d}][{d}] = {c}\n", .{ i, j, cell });
        }
        std.debug.print("\n", .{});
    }
}

pub fn main() void {
    std.debug.print("Hello, World!\n", .{});
    std.debug.print("Hello, World2!\n", .{});
}
