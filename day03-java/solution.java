import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Scanner;
import java.util.regex.MatchResult;
import java.util.regex.Pattern;
import java.util.stream.Collector;
import java.util.stream.Collectors;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;

public class solution {
    public static void main(String[] args) {

    
        test(sol1("test-input.txt"), 161);
        test(sol1("input.txt"), 161);
    }


    private static String readFile(String path) {
        try {
            return new String(Files.readAllBytes(Paths.get(path)));
        } catch (IOException e) {
            e.printStackTrace();
            return null;
        }
    }

    

    private static void test(int actual, int expected) {
        if (actual != expected) {

            System.out.println(String.format("❌ Test fail, Expected: %s, Actual: %s", expected, actual));
            
        } else {
            System.out.println("✅ Test pass");
            
        }
    }

    private static int sol1(String fileName) {
        var input = readFile(fileName);
        var sum = 0;
        var regex = "mul\\((\\d+),(\\d+)\\)";
        var sections = Pattern.compile(regex)
                .matcher(input)
                .results()
                .map(r -> {
                    var x = Integer.parseInt(r.group(1));
                    var y = Integer.parseInt(r.group(2));
                    return new Mul(x, y);
                })
                .collect(Collectors.toList());

        for (Mul mul : sections) {
            sum = sum + (mul.x * mul.y);
        }
        return sum;
    }

    public record Mul(int x, int y) {}
}
