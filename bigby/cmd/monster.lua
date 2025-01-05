function MonsterHelp()
	print("USAGE: bigby monster [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create monster")
	print("  update     Update monster")
	print("  get        Get monster")
	print("  delete     Delete monster")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function MonsterCreateHelp()
	print("Usage: bigby monster create [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function MonsterUpdateHelp()
	print("Usage: bigby monster update [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function MonsterGetHelp()
	print("Usage: bigby monster get [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function MonsterDeleteHelp()
	print("Usage: bigby monster delete [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function MonsterHandler(args)
	for i = 2, #args do
		if args[i] == "--help" then
			MonsterHelp()
			return
		end
	end
end
