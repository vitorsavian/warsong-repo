function CharacterHelp()
	print("Usage: owlbear character [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create character")
	print("  update     Update character")
	print("  get        Get character")
	print("  delete     Delete character")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterCreateHelp()
	print("Usage: owlbear character create [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterUpdateHelp()
	print("Usage: owlbear character update [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterGetHelp()
	print("Usage: owlbear character get [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterDeleteHelp()
	print("Usage: owlbear character delete [OPTIONS]")
	print("COMMAND:")
	print("")
	print("OPTIONS:")
	print("  --help     Show this message")
end

function CharacterHandler(args)
	local arguments = {}

	for i = 3, #args do
		if args[i] == "--help" then
			CharacterHelp()
			return
		end

		print("Getting all the things")

		if arg == "--name" then
			arguments.name = args[i + 1]
		elseif arg == "--level" then
			arguments.level = args[i + 1]
		elseif arg == "--str" then
			arguments.str = args[i + 1]
		elseif arg == "--dex" then
			arguments.dex = args[i + 1]
		elseif arg == "--con" then
			arguments.con = args[i + 1]
		elseif arg == "--int" then
			arguments.int = args[i + 1]
		elseif arg == "--char" then
			arguments.char = args[i + 1]
		elseif arg == "--wil" then
			arguments.wil = args[i + 1]
		end
	end

	return arguments
end
