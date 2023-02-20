
package com.example.demo;

import java.util.Map;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
@RestController
public class DemoApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemoApplication.class, args);
	}

	@PostMapping("/invoke")
	public String invoke(@RequestHeader Map<String, String> headers, @RequestBody String payload) {
		// 注意：JAVA 为编译型语言，直接修改代码不能直接生效！请在控制台右上角“导出代码”，然后根据 HELP.md 中的说明编译代码并重新上传。
		// 注意：JAVA 为编译型语言，直接修改代码不能直接生效！请在控制台右上角“导出代码”，然后根据 HELP.md 中的说明编译代码并重新上传。
		// 注意：JAVA 为编译型语言，直接修改代码不能直接生效！请在控制台右上角“导出代码”，然后根据 HELP.md 中的说明编译代码并重新上传。
		// Notice: You need to complie the code first otherwise the code change will not
		// take effect.
		String requestId = headers.get("x-fc-request-id");
		// System.out.printf("FC Invoke Start RequestId: %s%n", requestId);
		System.out.println(payload);
		// System.out.printf("FC Invoke End RequestId: %s%n", requestId);
		return payload;
	}
}