const std = @import("std");

fn readFileAndParse(comptime path: []const u8, allocator: std.mem.Allocator) !std.ArrayList(std.ArrayList(u8)) {
    const file = @embedFile(path);
    var it = std.mem.tokenize(u8, file, "\n");
    var grid = std.ArrayList(std.ArrayList(u8)).init(allocator);

    while (it.next()) |line| {
        var row = std.ArrayList(u8).init(allocator);
        try row.appendSlice(line);
        try grid.append(row);
    }

    return grid;
}

test "sol1" {
    var sum: i32 = 0;
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    const grid = try readFileAndParse("input.txt", allocator);

    defer {
        for (grid.items) |row| {
            row.deinit();
        }
        grid.deinit();
    }

    const row = grid.items.len;
    const column = grid.items[0].items.len;

    for (0..row) |i| {
        for (0..column) |j| {

            //std.debug.print("Sum: {d}\n", .{sum});
            // Horizontal
            if (j + 3 < column) {
                const horizontal = grid.items[i].items[j .. j + 4];

                std.debug.print("Horizontal: {s}\n", .{horizontal});
                if (std.mem.eql(u8, horizontal, "XMAS") or std.mem.eql(u8, horizontal, "SAMX")) {
                    std.debug.print("Found horizontal: {s}\n", .{horizontal});
                    sum += 1;
                }
            }

            // Vertical
            if (i + 3 < row) {
                const vertical = [_]u8{
                    grid.items[i].items[j],
                    grid.items[i + 1].items[j],
                    grid.items[i + 2].items[j],
                    grid.items[i + 3].items[j],
                };

                std.debug.print("Vertical: {s}\n", .{vertical});
                if (std.mem.eql(u8, &vertical, "XMAS") or std.mem.eql(u8, &vertical, "SAMX")) {
                    std.debug.print("Found vertical: {s}\n", .{vertical});
                    sum += 1;
                }
            }

            // Diagonal Left-to-Right
            if (i + 3 < row and j + 3 < column) {
                const diagonalLTR = [_]u8{
                    grid.items[i].items[j],
                    grid.items[i + 1].items[j + 1],
                    grid.items[i + 2].items[j + 2],
                    grid.items[i + 3].items[j + 3],
                };

                if (std.mem.eql(u8, &diagonalLTR, "XMAS") or std.mem.eql(u8, &diagonalLTR, "SAMX")) {
                    std.debug.print("Found diagonal Left-to-Right: {s}\n", .{diagonalLTR});
                    sum += 1;
                }
            }

            // Diagonal Right-to-Left
            const safeJ: i32 = @intCast(j);
            if (i + 3 < row and safeJ - 3 >= 0) {
                const diagonalRTL = [_]u8{
                    grid.items[i].items[j],
                    grid.items[i + 1].items[j - 1],
                    grid.items[i + 2].items[j - 2],
                    grid.items[i + 3].items[j - 3],
                };

                if (std.mem.eql(u8, &diagonalRTL, "XMAS") or std.mem.eql(u8, &diagonalRTL, "SAMX")) {
                    std.debug.print("Found diagonal Right-to-Left: {s}\n", .{diagonalRTL});
                    sum += 1;
                }
            }
        }
    }

    std.debug.print("Sum: {d}\n", .{sum});

    try std.testing.expectEqual(sum, 2406);
}

pub fn main() void {
    std.debug.print("Hello, World!\n", .{});
    std.debug.print("Hello, World2!\n", .{});
}
