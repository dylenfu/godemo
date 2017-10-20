package base

func Route(sub string) {
	switch sub {
	case "hi":
		println("Hello IPFS")

	case "simple-string":
		SimpleStringType()

	case "simple-atomic":
		SimpleAtomic()

	case "simple-struct":
		t := &Tee{3, "hello"}
		println(t.GetTeeNum())
		println(t.GetTeeName())

	case "simple-channel":
		ChannelDemo()

	case "simple-channel-wait":
		SimpleChannelDemo()

	case "simple-interface":
		InterfaceDemo()

	case "reflect-demo1":
		ReflectDemo1()

	case "reflect-demo2":
		ReflectDemo2()

	case "reflect-demo3":
		ReflectDemo3()

	case "reflect-demo4":
		ReflectDemo4()

	case "reflect-demo5":
		ReflectDemo5()

	case "reflect-simple-call":
		ReflectSimpleCall()

	case "reflect-simple-struct-call":
		ReflectStructCall()

	case "reflect-judge":
		JudgeType()

	case "simple-path":
		SimplePath()

	case "simple-os-args":
		SimpleOsArgs()
		break

	case "simple-math":
		SimpleMath()
	}
}