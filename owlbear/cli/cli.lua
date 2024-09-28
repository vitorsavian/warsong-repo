function MainHelp()
	print("USAGE: owlbear [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  character   character command")
	print("  npc         npc command")
	print("  monster     monster command")
	print("OPTIONS:")
	print("  --help      show this message")
end

-- TODO: rewrite the entire character and monster help to match the commands that I want
function CharacterHelp()
	print("Usage: owlbear character [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create character")
	print("  update     Update character")
	print("  get        Get character")
	print("  delete     Delete character")
end

function NpcHelp()
	print("Usage: owlbear npc [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create npc")
	print("  update     Update npc")
	print("  get        Get npc")
	print("  delete     Delete npc")
end

function MonsterHelp()
	print("USAGE: owlbear monster [COMMAND] [OPTIONS]")
	print("COMMAND:")
	print("  create     Create monster")
	print("  update     Update monster")
	print("  get        Get monster")
	print("  delete     Delete monster")
end

function Main(args)
	if #args < 1 then
		Print_main_help()
		return
	end

	if args[1] == "create" and args[2] == "character" then
		local name = nil

		for i = 3, #args do
			if args[i] == "--name" then
				if i + 1 <= #args then
					name = args[i + 1]
				else
					print("Erro: --name requer um valor.")
					return
				end
			elseif args[i] == "--help" then
				Print_help()
				return
			else
				print("Erro: opção desconhecida " .. args[i])
				return
			end
		end

		if name then
			print("Character and stuff")
			-- create_character(name)
		else
			print("Erro: Nome do personagem não fornecido.")
		end
	else
		Print_main_help()
	end
end

-- Call the main function
Main(arg)
