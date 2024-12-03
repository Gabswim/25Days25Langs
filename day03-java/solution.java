import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.regex.Pattern;
import java.io.IOException;

public class solution {
    public static void main(String[] args) {

        test(sol1("test-input.txt"), 161);
        test(sol1("input.txt"), 173529487);

        test(sol2("test-input-2.txt"), 48);
        test(sol2("input.txt"), 99532691);
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

        var regex = "mul\\((\\d+),(\\d+)\\)";

        var sum = Pattern.compile(regex)
                .matcher(input)
                .results()
                .map(r -> {
                    var x = Integer.parseInt(r.group(1));
                    var y = Integer.parseInt(r.group(2));
                    return new Mul(x, y);
                })
                .mapToInt(mul -> mul.x * mul.y)
                .reduce(0, Integer::sum);

        return sum;
    }

    private static int sol2(String fileName) {
        var input = readFile(fileName);

        var regex = "(don't\\(\\)|do\\(\\)|mul\\((\\d+),(\\d+)\\))";

        var result = (Result) Pattern.compile(regex)
                .matcher(input)
                .results()
                .map(r -> {
                    if (r.group(0).startsWith("mul(")) {
                        var x = Integer.parseInt(r.group(2));
                        var y = Integer.parseInt(r.group(3));
                        return new Mul(x, y);
                    }
                    if (r.group(0).startsWith("do()")) {
                        return new Do();
                    }
                    if (r.group(0).startsWith("don't()")) {
                        return new DoNot();
                    }
                    throw new Error("Something is not correct in the regex");
                }).reduce(new Result(true, 0), (acc, element) -> {

                    var _result = (Result) acc;

                    var enabled = _result.enabled;
                    var sum = _result.sum;
                    if (enabled && element instanceof Mul) {
                        var mul = (Mul) element;
                        return new Result(enabled, _result.sum + (mul.x * mul.y));
                    }
                    if (element instanceof Do) {
                        return new Result(true, sum);
                    }
                    if (element instanceof DoNot) {
                        return new Result(false, sum);
                    }

                    return new Result(enabled, sum);
                });

        return result.sum;
    }

    public record Mul(int x, int y) {
    }

    public record Do() {
    }

    public record DoNot() {
    }

    public record Result(boolean enabled, int sum) {
    }
}
